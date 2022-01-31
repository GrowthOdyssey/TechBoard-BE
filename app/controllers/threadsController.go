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
	switch r.Method {
	case http.MethodGet:
		page := r.FormValue("page")
		perPage := r.FormValue("perPage")
		threads := getThreads(page, perPage)
		json.NewEncoder(w).Encode(threads)
	case http.MethodPost:
		accessToken := r.Header.Get("accessToken")
		if accessToken == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		e := r.ParseForm()
		fmt.Println(e, accessToken)
		threadTitle := r.Form.Get("threadTitle")
		categoryId := r.Form.Get("categoryId")
		newThread := postThread(accessToken, threadTitle, categoryId)
		json.NewEncoder(w).Encode(newThread)
	default:
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

// スレッドハンドラ（パスパラメータが存在する場合）
func threadsIdHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/v1/threads/")
	switch r.Method {
	case http.MethodGet:
		thread := getThreadById(id)
		json.NewEncoder(w).Encode(thread)
	// MEMO URLにcomments入っているか判定してハンドリングしたいかも
	case http.MethodPost:
		e := r.ParseForm()
		fmt.Println(e)
		userId := r.Form.Get("userId")
		sessionId := r.Form.Get("sessionId")
		commentTitle := r.Form.Get("commentTitle")
		comment := postThreadComments(id, userId, sessionId, commentTitle)
		json.NewEncoder(w).Encode(comment)
	default:
		// TODO aiharanaoya
		// 仮で500のStatusTextを返している。今後エラーハンドリングを実装。
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

// コントローラ関数

// スレッド一覧取得
func getThreads(page, perPage string) *models.ThreadsAndPagination {
	fmt.Println("スレッド一覧取得処理")
	return models.GetThreadsSql(page, perPage)
}

// スレッド作成
func postThread(accessToken, threadTitle, categoryId string) *models.ThreadAndUser {
	fmt.Println("スレッド作成処理")
	return models.PostThreadSql(accessToken, threadTitle, categoryId)
}

// スレッド取得
func getThreadById(id string) *models.ThreadAndComments {
	fmt.Println("スレッド取得処理")
	return models.GetThreadByIdSql(id)
}

// スレッドコメント作成
func postThreadComments(id, userId, sessionId, commentTitle string) *models.CommentAndThreadAndUser {
	fmt.Println("スレッドコメント作成処理")
	return models.PostCommentsSql(id, userId, sessionId, commentTitle)
}
