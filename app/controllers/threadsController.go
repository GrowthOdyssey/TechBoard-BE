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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(ErrMsg{"パスにIDがありません"})
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	case regexp.MustCompile(`[^0-9]`).Match([]byte(id)):
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(ErrMsg{"パスは数字で指定してください"})
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
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
	var errMsgAndErrors struct {
		ErrMessage string `json:"message"`
		Errors     struct {
			CategoryId string `json:"categoryId"`
			Page       string `json:"page"`
			PerPage    string `json:"perPage"`
		} `json:"errors"`
	}
	categoryId := r.FormValue("categoryId")
	if regexp.MustCompile(`[^0-9]`).Match([]byte(categoryId)) {
		errMsgAndErrors.Errors.CategoryId = "categoryIdは数字で指定してください"
	}
	page := r.FormValue("page")
	if page == "" {
		errMsgAndErrors.Errors.Page = "pageを指定してください"
	} else if regexp.MustCompile(`[^0-9]`).Match([]byte(page)) {
		errMsgAndErrors.Errors.Page = "pageは数字で指定してください"
	}
	perPage := r.FormValue("perPage")
	if perPage == "" {
		errMsgAndErrors.Errors.PerPage = "perPageを指定してください"
	} else if regexp.MustCompile(`[^0-9]`).Match([]byte(perPage)) {
		errMsgAndErrors.Errors.PerPage = "perPageは数字で指定してください"
	}

	if errMsgAndErrors.Errors.CategoryId+errMsgAndErrors.Errors.Page+errMsgAndErrors.Errors.PerPage != "" {
		errMsgAndErrors.ErrMessage = "値が不正です"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(422)
		json.NewEncoder(w).Encode(errMsgAndErrors)
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
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
	var errMsgAndErrors struct {
		ErrMessage string `json:"message"`
		Errors     struct {
			AccessToken string `json:"accessToken"`
			ThreadTitle string `json:"threadTitle"`
			CategoryId  string `json:"categoryId"`
		} `json:"errors"`
	}
	accessToken := r.Header.Get("accessToken")
	if accessToken == "" {
		errMsgAndErrors.Errors.AccessToken = "accessTokenがありません"
	}

	var reqBody struct {
		ThreadTitle string `json:"threadTitle"`
		CategoryId  string `json:"categoryId"`
	}
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		fmt.Println(err)
	}
	if reqBody.ThreadTitle == "" {
		errMsgAndErrors.Errors.ThreadTitle = "threadTitleを入力してください"
	}
	if reqBody.CategoryId == "" {
		errMsgAndErrors.Errors.CategoryId = "categoryIdを選択してください"
	} else if regexp.MustCompile(`[^0-9]`).Match([]byte(reqBody.CategoryId)) {
		errMsgAndErrors.Errors.CategoryId = "categoryIdは数字で指定してください"
	}
	if errMsgAndErrors.Errors.AccessToken+errMsgAndErrors.Errors.ThreadTitle+errMsgAndErrors.Errors.CategoryId != "" {
		errMsgAndErrors.ErrMessage = "値が不正です"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(422)
		json.NewEncoder(w).Encode(errMsgAndErrors)
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	newThreadId := models.PostThreadSql(accessToken, reqBody.ThreadTitle, reqBody.CategoryId)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
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
	var errMsgAndErrors struct {
		ErrMessage string `json:"message"`
		Errors     struct {
			UserId       string `json:"userId"`
			SessionId    string `json:"sessionId"`
			CommentTitle string `json:"commentTitle"`
		} `json:"errors"`
	}
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
		errMsgAndErrors.Errors.UserId = "userIdを入力してください"
		errMsgAndErrors.Errors.SessionId = "sessionIdを入力してください"
	}
	if reqBody.CommentTitle == "" {
		errMsgAndErrors.Errors.CommentTitle = "commentTitleを入力してください"
	}
	if errMsgAndErrors.Errors.UserId+errMsgAndErrors.Errors.SessionId+errMsgAndErrors.Errors.CommentTitle != "" {
		errMsgAndErrors.ErrMessage = "値が不正です"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(422)
		json.NewEncoder(w).Encode(errMsgAndErrors)
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	comment, modelErr := models.PostCommentsSql(threadId, reqBody.UserId, reqBody.SessionId, reqBody.CommentTitle)
	if modelErr != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(modelErr)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
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
