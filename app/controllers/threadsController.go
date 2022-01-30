package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GrowthOdyssey/TechBoard-BE/app/models"
)

// スレッドハンドラ
func threadsHandler(w http.ResponseWriter, r *http.Request /* ,page,perPage int */) {
	switch r.Method {
	case http.MethodGet:
		threads := getThreads(1, 2) //今はpage,perPageベタ打ち
		json.NewEncoder(w).Encode(threads)
	case http.MethodPost:
		postThread()
	default:
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

// スレッドハンドラ（パスパラメータが存在する場合）
func threadsIdHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getThreadById()
	// MEMO URLにcomments入っているか判定してハンドリングしたいかも
	case http.MethodPost:
		postThreadComments()
	default:
		// TODO aiharanaoya
		// 仮で500のStatusTextを返している。今後エラーハンドリングを実装。
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

// コントローラ関数

// スレッド一覧取得
func getThreads(page, perPage int) *models.ThreadsAndPagination {
	fmt.Println("スレッド一覧取得処理")
	return models.GetThreadsSql(page, perPage)
}

// スレッド作成
func postThread() {
	fmt.Println("スレッド作成処理")
}

// スレッド取得
func getThreadById() {
	fmt.Println("スレッド取得処理")
}

// スレッドコメント作成
func postThreadComments() {
	fmt.Println("スレッドコメント作成処理")
}
