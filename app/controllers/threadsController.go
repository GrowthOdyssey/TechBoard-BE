package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/GrowthOdyssey/TechBoard-BE/app/models"
)

type ErrMsg struct {
	ErrorMessage string `json:"message"`
}

type ErrMsgAndErrors struct {
	ErrMessage string   `json:"message"`
	Errors     []string `json:"errors"`
}

func err400(w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)
	json.NewEncoder(w).Encode(ErrMsg{msg})
}

// スレッドハンドラ
func threadsHandler(w http.ResponseWriter, r *http.Request) {
	allowCors(w)
	switch r.Method {
	case http.MethodGet:
		getThreads(w, r)
	case http.MethodPost:
		postThread(w, r)
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
	default:
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

// スレッドハンドラ（パスパラメータが存在する場合）
func threadsIdHandler(w http.ResponseWriter, r *http.Request) {
	allowCors(w)
	id := strings.TrimSuffix(strings.TrimPrefix(r.URL.Path, "/v1/threads/"), "/comments")
	switch {
	case id == "":
		err400(w, "パスにIDがありません")
	case regexp.MustCompile(`[^0-9]`).Match([]byte(id)):
		err400(w, "パスは数字で指定してください")
	default:
		switch r.Method {
		case http.MethodGet:
			getThreadById(w, r, id)
		// MEMO URLにcomments入っているか判定してハンドリングしたいかも
		case http.MethodPost:
			postThreadComments(w, r, id)
		case http.MethodOptions:
			w.WriteHeader(http.StatusOK)
		default:
			// TODO aiharanaoya
			// 仮で500のStatusTextを返している。今後エラーハンドリングを実装。
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}

// スレッドカテゴリーハンドラ（パスパラメータが存在する場合）
func threadsCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	allowCors(w)
	switch r.Method {
	case http.MethodGet:
		getThreadsCategories(w, r)
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
	default:
		// TODO aiharanaoya
		// 仮で500のStatusTextを返している。今後エラーハンドリングを実装。
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

// コントローラ関数

// スレッド一覧取得
func getThreads(w http.ResponseWriter, r *http.Request) {
	fmt.Println("スレッド一覧取得処理")
	categoryId := r.FormValue("categoryId")
	if regexp.MustCompile(`[^0-9]`).Match([]byte(categoryId)) {
		err400(w, "categoryIdは数字で指定してください")
		return
	}
	page := r.FormValue("page")
	if page == "" {
		err400(w, "pageを指定してください")
		return
	} else if regexp.MustCompile(`[^0-9]`).Match([]byte(page)) {
		err400(w, "pageは数字で指定してください")
		return
	}
	perPage := r.FormValue("perPage")
	if perPage == "" {
		err400(w, "perPageを指定してください")
		return
	} else if regexp.MustCompile(`[^0-9]`).Match([]byte(perPage)) {
		err400(w, "perPageは数字で指定してください")
		return
	}
	threads := models.GetThreadsSql(categoryId, page, perPage)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(threads)
}

// スレッド作成
func postThread(w http.ResponseWriter, r *http.Request) {
	fmt.Println("スレッド作成処理")

	accessToken := r.Header.Get("accessToken")
	if accessToken == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(401)
		json.NewEncoder(w).Encode(ErrMsg{"accessTokenがありません"})
		return
	}

	isOk := AuthorizationCheck(w, accessToken)
	if !isOk {
		return
	}

	errMsgAndErrors := ErrMsgAndErrors{}
	var reqBody struct {
		ThreadTitle string `json:"threadTitle"`
		CategoryId  string `json:"categoryId"`
	}
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		fmt.Println(err)
	}

	if reqBody.ThreadTitle == "" {
		errMsgAndErrors.Errors = append(errMsgAndErrors.Errors, "スレッド名を入力してください")
	}
	if reqBody.CategoryId == "" {
		errMsgAndErrors.Errors = append(errMsgAndErrors.Errors, "カテゴリーを選択してください")
	} else if regexp.MustCompile(`[^0-9]`).Match([]byte(reqBody.CategoryId)) {
		errMsgAndErrors.Errors = append(errMsgAndErrors.Errors, "categoryIdは数字で指定してください")
	}
	if len(errMsgAndErrors.Errors) != 0 {
		errMsgAndErrors.ErrMessage = "値が不正です"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(422)
		json.NewEncoder(w).Encode(errMsgAndErrors)
		return
	}

	newThreadId := models.PostThreadSql(accessToken, reqBody.ThreadTitle, reqBody.CategoryId)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(newThreadId)
}

// スレッド取得
func getThreadById(w http.ResponseWriter, r *http.Request, id string) {
	fmt.Println("スレッド取得処理")
	thread := models.GetThreadByIdSql(id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(thread)
}

// スレッドコメント作成
func postThreadComments(w http.ResponseWriter, r *http.Request, threadId string) {
	fmt.Println("スレッドコメント作成処理")
	errMsgAndErrors := ErrMsgAndErrors{}
	var reqBody struct {
		UserId       string `json:"userId"`
		SessionId    string `json:"sessionId"`
		CommentTitle string `json:"commentTitle"`
	}
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		fmt.Println(err)
	}
	if reqBody.UserId == "" && reqBody.SessionId == "" {
		err400(w, "userIdとsessionIdどちらもありません")
		return
	}
	if reqBody.CommentTitle == "" {
		errMsgAndErrors.Errors = append(errMsgAndErrors.Errors, "コメントを入力してください")
		errMsgAndErrors.ErrMessage = "値が不正です"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(422)
		json.NewEncoder(w).Encode(errMsgAndErrors)
		return
	}

	comment, modelErr := models.PostCommentsSql(threadId, reqBody.UserId, reqBody.SessionId, reqBody.CommentTitle)
	if modelErr != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(modelErr)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(comment)
}

// スレッドカテゴリー一覧取得
func getThreadsCategories(w http.ResponseWriter, r *http.Request) {
	fmt.Println("スレッドカテゴリー一覧取得処理")
	//id := r.FormValue("id")
	categories := models.GetThreadsCategoriesSql()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(categories)
}
