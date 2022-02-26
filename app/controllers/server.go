package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/GrowthOdyssey/TechBoard-BE/app/constants"
	"github.com/GrowthOdyssey/TechBoard-BE/config"
)

// ヘルスチェック構造体
type HealthCheck struct {
	Status string `json:"status"`
}

// 共通エラー構造体
type CommonError struct {
	Message string `json:"message"`
}

// ルーティング設定
func SetRouter() {
	// ヘルスチェック
	http.HandleFunc("/health_check", healthCheck)

	// ユーザー
	http.HandleFunc("/v1/users", usersHandler)
	http.HandleFunc("/v1/users/login", usersLoginHandler)
	http.HandleFunc("/v1/users/logout", usersLogoutHandler)

	// 記事
	http.HandleFunc("/v1/articles", articlesHandler)
	http.HandleFunc("/v1/articles/", articlesIdHandler)

	// スレッド
	http.HandleFunc("/v1/threads", threadsHandler)
	http.HandleFunc("/v1/threads/", threadsIdHandler)
	http.HandleFunc("/v1/threads/categories", threadsCategoriesHandler)
}

// サーバーを起動する
func StartServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = config.Config.Port
	}
	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		fmt.Println(err)
	}
}

// ヘルスチェック
func healthCheck(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		healthCheck := HealthCheck{Status: "OK"}

		healthCheckRes, err := json.Marshal(healthCheck)
		if err != nil {
			fmt.Println(err)
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(healthCheckRes))
	} else {
		ResponseCommonError(w, http.StatusNotFound, constants.NotFoundMessage)
	}
}

// CORS許可
func allowCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
}

// 共通エラーレスポンス
func ResponseCommonError(w http.ResponseWriter, statusCode int, message string) {
	error := CommonError{Message: message}

	errorRes, err := json.Marshal(error)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	fmt.Fprint(w, string(errorRes))
}
