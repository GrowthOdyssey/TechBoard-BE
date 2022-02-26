package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GrowthOdyssey/TechBoard-BE/app/constants"
	"github.com/GrowthOdyssey/TechBoard-BE/app/models"
)

// ハンドラ関数
// URL、HTTPメソッドから呼び出す関数をハンドリングする。
// 基本的にコントローラ関数を呼び出すのみで処理はコントローラ関数に記載する。

// ユーザーハンドラ
func usersHandler(w http.ResponseWriter, r *http.Request) {
	allowCors(w)
	switch r.Method {
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
	case http.MethodGet:
		getUser(w, r)
	case http.MethodPost:
		usersSignUp(w, r)
	default:
		ResponseCommonError(w, http.StatusNotFound, constants.NotFoundMessage)
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
		ResponseCommonError(w, http.StatusNotFound, constants.NotFoundMessage)
	}
}

// ログアウトハンドラ
func usersLogoutHandler(w http.ResponseWriter, r *http.Request) {
	allowCors(w)
	switch r.Method {
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
	case http.MethodDelete:
		usersLogout(w, r)
	default:
		ResponseCommonError(w, http.StatusNotFound, constants.NotFoundMessage)
	}
}

// コントローラ関数
// それぞれのAPIに対応した関数。
// モデル関数で定義した構造体の呼び出し、JSONの変換処理等を行う。
// DBのアクセス関数、レシーバメソッド、複雑になるロジックはモデル関数に定義する。

func getUser(w http.ResponseWriter, r *http.Request) {
	accessToken := r.Header.Get("accessToken")
	if accessToken == "" {
		ResponseCommonError(w, http.StatusNotFound, constants.UnauthorizedMessage)
	}

	AuthorizationCheck(w, accessToken)

	userRes, err := models.GetUser(accessToken)
	if err != nil {
		fmt.Println(err)
	}

	// JSON変換
	userResJson, err := json.Marshal(userRes)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(userResJson))
}

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

	// ログイン失敗時は401エラー
	if !isOk {
		ResponseCommonError(w, http.StatusUnauthorized, constants.LoginErrorMessage)
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
		ResponseCommonError(w, http.StatusNotFound, constants.UnauthorizedMessage)
	}

	AuthorizationCheck(w, accessToken)

	err := 	models.Logout(accessToken)
	if err != nil {
		fmt.Println(err)
	}

	w.WriteHeader(http.StatusNoContent)
}

// アクセストークンからログインされているかチェックする
// ログインされていなかったら401エラーを返す
func AuthorizationCheck(w http.ResponseWriter, accessToken string) {
	isOk, err := models.CheckLoginByAccessToken(accessToken)
	if err != nil {
		fmt.Println(err)
	}

	if err != nil || !isOk {
		ResponseCommonError(w, http.StatusUnauthorized, constants.UnauthorizedMessage)
	}
}
