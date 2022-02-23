package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/GrowthOdyssey/TechBoard-BE/app/models"
)

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
	page := r.FormValue("page")
	perPage := r.FormValue("perPage")
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
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	var reqBody struct {
		ThreadTitle string `json:"threadTitle"`
		CategoryId  string `json:"categoryId"`
	}
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		fmt.Println(err)
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
	var reqBody struct {
		UserId       string `json:"userId"`
		SessionId    string `json:"sessionId"`
		CommentTitle string `json:"commentTitle"`
	}
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		fmt.Println(err)
	}
	comment := models.PostCommentsSql(threadId, reqBody.UserId, reqBody.SessionId, reqBody.CommentTitle)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(comment)
}

// スレッドコメント作成
func getThreadsCategories(w http.ResponseWriter, r *http.Request) {
	fmt.Println("スレッドカテゴリー一覧取得処理")
	categories := models.GetThreadsCategoriesSql()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(categories)
}
