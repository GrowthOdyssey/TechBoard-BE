package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GrowthOdyssey/TechBoard-BE/app/models"
)

// ハンドラ関数
// URL、HTTPメソッドから呼び出す関数をハンドリングする。
// 基本的にコントローラ関数を呼び出すのみで処理はコントローラ関数に記載する。

// ユーザー登録ハンドラ
func usersSignUpHandler(w http.ResponseWriter, r *http.Request) {
	allowCors(w)
	switch r.Method {
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
	case http.MethodPost:
		usersSignUp(w, r)
	default:
		// TODO aiharanaoya
		// 仮で500のStatusTextを返している。今後エラーハンドリングを実装。
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

// ログインハンドラ
func usersLoginHandler(w http.ResponseWriter, r *http.Request) {
	allowCors(w)
	switch r.Method {
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
	case http.MethodPost:
		usersLogin(w, r)
	default:
		// TODO aiharanaoya
		// 仮で500のStatusTextを返している。今後エラーハンドリングを実装。
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

// ログアウトハンドラ
func usersLogoutHandler(w http.ResponseWriter, r *http.Request) {
	allowCors(w)StatusOK
	switch r.Method {
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
	case http.MethodDelete:
		usersLogout(w, r)
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

// ユーザー登録
func usersSignUp(w http.ResponseWriter, r *http.Request) {
	userSignUpReq := models.UserSignUpReq{}

	err := json.NewDecoder(r.Body).Decode(&userSignUpReq)
	if err != nil {
		fmt.Println(err)
	}

	// ユーザー登録する
	userRes, err := userSignUpReq.RegisterUser()
	if err != nil {
		fmt.Println(err)
	}

	// ユーザー登録後、そのままログインする
	accessToken, err := models.Login(userRes.UserId)
	if err != nil {
		fmt.Println(err)
	}

	// ログイン時生成したアクセストークンをレスポンスに加える
	userRes.AccessToken = accessToken

	// JSON変換
	userResJson, err := json.Marshal(userRes)
	if err != nil {
			fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, string(userResJson))
}

// ログイン
func usersLogin(w http.ResponseWriter, r *http.Request) {
	userLoginReq := models.UserLoginReq{}

	err := json.NewDecoder(r.Body).Decode(&userLoginReq)
	if err != nil {
		fmt.Println(err)
	}

	isOk, userRes, err := userLoginReq.CheckLogin()
	if err != nil {
		fmt.Println(err)
	}

	if !isOk {
		// TODO aiharanaoya 仮
		http.Error(w, "ログイン失敗", http.StatusUnauthorized)
		return
	}

	accessToken, err := models.Login(userLoginReq.UserId)
	if err != nil {
		fmt.Println(err)
	}

	// ログイン時生成したアクセストークンをレスポンスに加える
	userRes.AccessToken = accessToken

	// JSON変換
	userResJson, err := json.Marshal(userRes)
	if err != nil {
			fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, string(userResJson))
}

// ログアウト
func usersLogout(w http.ResponseWriter, r *http.Request) {
	accessToken := r.Header.Get("accessToken")
	if accessToken == "" {
		// TODO aiharanaoya
		// 仮で500のStatusTextを返している。今後エラーハンドリングを実装。
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	err := 	models.Logout(accessToken)
	if err != nil {
		fmt.Println(err)
	}

	w.WriteHeader(http.StatusNoContent)
}
