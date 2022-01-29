package controllers

import (
	"fmt"
	"net/http"
)

// ユーザー登録ハンドラ
func usersSignUpHandler(w http.ResponseWriter, r *http.Request) {
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
	if r.Method == http.MethodDelete {
		usersLogout()
	} else {
		// TODO aiharanaoya
		// 仮で500のStatusTextを返している。今後エラーハンドリングを実装。
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

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