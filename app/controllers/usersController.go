package controllers

import (
	"fmt"
	"net/http"
)

// ハンドラ関数
// URL、HTTPメソッドから呼び出す関数をハンドリングする。
// 基本的にコントローラ関数を呼び出すのみで処理はコントローラ関数に記載する。

// ユーザー登録ハンドラ
func usersSignUpHandler(w http.ResponseWriter, r *http.Request) {
	allowCors(w)
	if r.Method == http.MethodPost {
		usersSignUp()
	} else {
		// TODO aiharanaoya
		// 仮で500のStatusTextを返している。今後エラーハンドリングを実装。
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

// ログインハンドラ
func usersLoginHandler(w http.ResponseWriter, r *http.Request) {
	allowCors(w)
	if r.Method == http.MethodPost {
		usersLogin()
	} else {
		// TODO aiharanaoya
		// 仮で500のStatusTextを返している。今後エラーハンドリングを実装。
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

// ログアウトハンドラ
func usersLogoutHandler(w http.ResponseWriter, r *http.Request) {
	allowCors(w)
	if r.Method == http.MethodDelete {
		usersLogout()
	} else {
		// TODO aiharanaoya
		// 仮で500のStatusTextを返している。今後エラーハンドリングを実装。
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

// コントローラ関数
// それぞれのAPIに対応した関数。
// モデル関数で定義した構造体の呼び出し、JSONの変換処理等を行う。
// DBのアクセス関数、レシーバメソッド、複雑になるロジックはモデル関数に定義する。

// ユーザー登録
func usersSignUp() {
	fmt.Println("ユーザー登録処理")
}

// ログイン
func usersLogin() {
	fmt.Println("ログイン処理")
}

// ログアウト
func usersLogout() {
	fmt.Println("ログアウト処理")
}