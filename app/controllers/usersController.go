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

// ユーザーバリデーションエラー構造体
type UserValidErrors struct {
	UserId string `json:"userId"`
	Password string `json:"password"`
	UserName string `json:"userName"`
}

// 422エラー構造体
type User422Error struct {
	Message string `json:"message"`
	Errors UserValidErrors `json:"errors"`
}

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
		return
	}

	isOk := AuthorizationCheck(w, accessToken)
	if !isOk {
		return
	}

	userRes, err := models.GetUser(accessToken)
	if err != nil {
		fmt.Println(err)
		ResponseCommonError(w, http.StatusInternalServerError, constants.InternalServerErrorMessage)
		return
	}

	// JSON変換
	userResJson, err := json.Marshal(userRes)
	if err != nil {
		fmt.Println(err)
		ResponseCommonError(w, http.StatusInternalServerError, constants.InternalServerErrorMessage)
		return
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
		ResponseCommonError(w, http.StatusInternalServerError, constants.InternalServerErrorMessage)
		return
	}

	var userIdValidMessage string
	var passwordValidMessage string
	var userNameValidMessage string

	// ユーザーIDバリデーション
	if userSignUpReq.UserId == "" {
		userIdValidMessage = constants.UserIdRequiredMessage
	}
	if len(userSignUpReq.UserId) > 20 {
		userIdValidMessage = constants.UserIdRequiredMessage
	}

	// パスワードバリデーション
	if userSignUpReq.Password == "" {
		passwordValidMessage = constants.PasswordRequiredMessage
	}
	if len(userSignUpReq.Password) < 8 {
		passwordValidMessage = constants.PasswordMinMessage
	}
	if len(userSignUpReq.Password) > 100 {
		passwordValidMessage = constants.PasswordMaxMessage
	}

	// ユーザー名バリデーション
	if userSignUpReq.UserName == "" {
		userNameValidMessage = constants.UserNameRequiredMessage
	}
	if len(userSignUpReq.UserName) > 100 {
		userNameValidMessage = constants.UserNameMaxMessage
	}

	if userIdValidMessage != "" || passwordValidMessage != "" || userNameValidMessage != "" {
		user422Error := User422Error{
			Message: constants.UnprocessableEntityErrorMessage,
			Errors: UserValidErrors{
				UserId: userIdValidMessage,
				Password: passwordValidMessage,
				UserName: userNameValidMessage,
			},
		}

		user422ErrorRes, err := json.Marshal(user422Error)
		if err != nil {
			fmt.Println(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprint(w, string(user422ErrorRes))
		return
	}

	// ユーザー登録する
	userRes, err := userSignUpReq.RegisterUser()
	if err != nil {
		fmt.Println(err)
		ResponseCommonError(w, http.StatusBadRequest, constants.ExistUser)
		return
	}

	// ユーザー登録後、そのままログインする
	accessToken, err := models.Login(userRes.UserId)
	if err != nil {
		fmt.Println(err)
		ResponseCommonError(w, http.StatusInternalServerError, constants.InternalServerErrorMessage)
		return
	}

	// ログイン時生成したアクセストークンをレスポンスに加える
	userRes.AccessToken = accessToken

	// JSON変換
	userResJson, err := json.Marshal(userRes)
	if err != nil {
		fmt.Println(err)
		ResponseCommonError(w, http.StatusInternalServerError, constants.InternalServerErrorMessage)
		return
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
		ResponseCommonError(w, http.StatusInternalServerError, constants.InternalServerErrorMessage)
		return
	}

	isOk, userRes, err := userLoginReq.CheckLogin()
	// ログイン失敗時は401エラー
	if err != nil || !isOk {
		fmt.Println(err)
		ResponseCommonError(w, http.StatusUnauthorized, constants.LoginErrorMessage)
		return
	}

	accessToken, err := models.Login(userLoginReq.UserId)
	if err != nil {
		fmt.Println(err)
		ResponseCommonError(w, http.StatusInternalServerError, constants.InternalServerErrorMessage)
		return
	}

	// ログイン時生成したアクセストークンをレスポンスに加える
	userRes.AccessToken = accessToken

	// JSON変換
	userResJson, err := json.Marshal(userRes)
	if err != nil {
		fmt.Println(err)
		ResponseCommonError(w, http.StatusInternalServerError, constants.InternalServerErrorMessage)
		return
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

	isOk := AuthorizationCheck(w, accessToken)
	if !isOk {
		return
	}

	err := 	models.Logout(accessToken)
	if err != nil {
		fmt.Println(err)
		ResponseCommonError(w, http.StatusInternalServerError, constants.InternalServerErrorMessage)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// アクセストークンからログインされているかチェックする
// ログインされていなかったら401エラーを返す
func AuthorizationCheck(w http.ResponseWriter, accessToken string) (isOk bool) {
	isOk, err := models.CheckLoginByAccessToken(accessToken)
	if err != nil {
		fmt.Println(err)
	}

	if err != nil || !isOk {
		ResponseCommonError(w, http.StatusUnauthorized, constants.UnauthorizedMessage)
	}

	return isOk
}
