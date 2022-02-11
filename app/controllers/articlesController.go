package controllers

import (
	"fmt"
	"net/http"
)

// ハンドラ関数
// URL、HTTPメソッドから呼び出す関数をハンドリングする。
// 基本的にコントローラ関数を呼び出すのみで処理はコントローラ関数に記載する。

// 記事ハンドラ
func articlesHandler(w http.ResponseWriter, r *http.Request) {
	allowCors(w)
	switch r.Method {
	case http.MethodGet:
		getArticles()
	case http.MethodPost:
		postArticle()
	default:
		// TODO aiharanaoya
		// 仮で500のStatusTextを返している。今後エラーハンドリングを実装。
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

// 記事ハンドラ（パスパラメータが存在する場合）
func articlesIdHandler(w http.ResponseWriter, r *http.Request) {
	allowCors(w)
	switch r.Method {
	case http.MethodGet:
		getArticleById()
	case http.MethodPut:
		putArticleById()
	case http.MethodDelete:
		deleteArticleById()
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

// 記事一覧取得
func getArticles() {
	fmt.Println("記事一覧取得処理")
}

// 記事作成
func postArticle() {
	fmt.Println("記事作成処理")
}

// 記事取得
func getArticleById() {
	fmt.Println("記事取得処理")
}

// 記事更新
func putArticleById() {
	fmt.Println("記事更新処理")
}

// 記事削除
func deleteArticleById() {
	fmt.Println("記事削除処理")
}