package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/GrowthOdyssey/TechBoard-BE/app/models"
)

// ハンドラ関数
// URL、HTTPメソッドから呼び出す関数をハンドリングする。
// 基本的にコントローラ関数を呼び出すのみで処理はコントローラ関数に記載する。

// 記事ハンドラ
func articlesHandler(w http.ResponseWriter, r *http.Request) {
	allowCors(w)
	switch r.Method {
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
	case http.MethodGet:
		getArticles(w,r)
	case http.MethodPost:
		postArticle(w,r)

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
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
	case http.MethodGet:
		getArticleById(w,r)
	case http.MethodPut:
		putArticleById(w,r)
	case http.MethodDelete:
		deleteArticleById(w,r)
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
func getArticles(w http.ResponseWriter, r *http.Request){
	userId := r.FormValue("userId")
	page := r.FormValue("page")
	perPage := r.FormValue("perPage")
	articles := models.GetArticlesSql(userId, page, perPage)
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}

// 記事作成
func postArticle(w http.ResponseWriter, r *http.Request) {
	accessToken := r.Header.Get("accessToken")
	//リクエストボディを取る準備
	var reqBody struct {
		ArticleTitle  string `json:"articleTitle"`
		Content  string `json:"content"`
	}
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		fmt.Println(err)
	}
//MODELに渡すための記述↓
	articles := models.PostArticleSql(accessToken,reqBody.ArticleTitle,reqBody.Content)
//MODELに渡すための記述↑

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}

// 記事取得
func getArticleById(w http.ResponseWriter, r *http.Request) {
	articleId := strings.TrimPrefix(r.URL.Path, "/v1/articles/")
	articles := models.GetArticleByIdSql(articleId)
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}

// 記事更新
func putArticleById(w http.ResponseWriter, r *http.Request) {
	accessToken := r.Header.Get("accessToken")
		//リクエストボディを取る準備
		var reqBody struct {
			ArticleId  string `json:"articleId"`
			ArticleTitle  string `json:"articleTitle"`
			Content  string `json:"content"`
		}
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		if err != nil {
			fmt.Println(err)
		}
	//MODELに渡すための記述↓
	articles := models.PutArticleSql(accessToken,reqBody.ArticleId,reqBody.ArticleTitle,reqBody.Content)
	//MODELに渡すための記述↑

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}

// 記事削除
func deleteArticleById(w http.ResponseWriter, r *http.Request) {
	accessToken := r.Header.Get("accessToken")
	articleId := strings.TrimPrefix(r.URL.Path, "/v1/articles/")
	models.DeleteArticleSql(accessToken,articleId)
	w.WriteHeader(204)
	w.Header().Set("Content-Type", "application/json")
}