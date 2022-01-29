package controllers

import (
	"fmt"
	"net/http"
)

// ハンドラ関数
// URL、HTTPメソッドから呼び出す関数をハンドリングする。
// 基本的にコントローラ関数を呼び出すのみで処理はコントローラ関数に記載する。

// スレッドハンドラ
func threadsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getThreads()
	case http.MethodPost:
		postThread()
	default:
		// TODO aiharanaoya
		// 仮で500のStatusTextを返している。今後エラーハンドリングを実装。
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
// それぞれのAPIに対応した関数。
// モデル関数で定義した構造体の呼び出し、JSONの変換処理等を行う。
// DBのアクセス関数、レシーバメソッド、複雑になるロジックはモデル関数に定義する。

// スレッド一覧取得
func getThreads() {
	fmt.Println("スレッド一覧取得処理")
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