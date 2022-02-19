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
		categoryId := r.FormValue("categoryId")
		page := r.FormValue("page")
		perPage := r.FormValue("perPage")
		threads := getThreads(categoryId, page, perPage)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(threads)
	case http.MethodPost:
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

		newThreadId := postThread(accessToken, reqBody.ThreadTitle, reqBody.CategoryId)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(newThreadId)
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
		thread := getThreadById(id)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(thread)
	// MEMO URLにcomments入っているか判定してハンドリングしたいかも
	case http.MethodPost:
		var reqBody struct {
			UserId       string `json:"userId"`
			SessionId    string `json:"sessionId"`
			CommentTitle string `json:"commentTitle"`
		}
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		if err != nil {
			fmt.Println(err)
		}
		comment := postThreadComments(id, reqBody.UserId, reqBody.SessionId, reqBody.CommentTitle)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(comment)
	default:
		// TODO aiharanaoya
		// 仮で500のStatusTextを返している。今後エラーハンドリングを実装。
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

// スレッドカテゴリーハンドラ（パスパラメータが存在する場合）
func threadsCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	allowCors(w)
	if r.Method == http.MethodGet {
		categories := getThreadsCategories()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(categories)
	} else {
		// TODO aiharanaoya
		// 仮で500のStatusTextを返している。今後エラーハンドリングを実装。
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

// コントローラ関数

// スレッド一覧取得
func getThreads(categoryId, page, perPage string) *models.ThreadsAndPagination {
	fmt.Println("スレッド一覧取得処理")
	return models.GetThreadsSql(categoryId, page, perPage)
}

// スレッド作成
func postThread(accessToken, threadTitle, categoryId string) *models.NewThreadId {
	fmt.Println("スレッド作成処理")
	return models.PostThreadSql(accessToken, threadTitle, categoryId)
}

// スレッド取得
func getThreadById(id string) *models.ThreadAndComments {
	fmt.Println("スレッド取得処理")
	return models.GetThreadByIdSql(id)
}

// スレッドコメント作成
func postThreadComments(id, userId, sessionId, commentTitle string) *models.CommentAndUser {
	fmt.Println("スレッドコメント作成処理")
	return models.PostCommentsSql(id, userId, sessionId, commentTitle)
}

// スレッドコメント作成
func getThreadsCategories() *models.ThreadsCategories {
	fmt.Println("スレッドカテゴリー一覧取得処理")
	return models.GetThreadsCategoriesSql()
}
