package models

import (
	"log"
	"strconv"
	"time"
	"crypto/rand"
	"errors"
)

// 記事一覧取得用ストラクト①
type Article struct {
	ArticleId               string    `json:"articleId"`
	ArticleTitle     string     `json:"articleTitle"`
	UserName     string     `json:"userName"`
	AvatarId     string     `json:"avatarId"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}
//記事一覧取得用ストラクト②
type ArticlesAndPagination struct {
	Articles    []Article   `json:"Articles"`//記事一覧取得用ストラクト①
	Pagination Pagination `json:"pagination"`
}

// 記事取得用ストラクト
type ArticleDetail struct {
	ArticleId               string    `json:"articleId"`
	ArticleTitle     string     `json:"articleTitle"`
	Content     string     `json:"content"`
	UserName     string     `json:"userName"`
	AvatarId     string     `json:"avatarId"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}
// 記事作成＆更新用ストラクト
type ArticleEdit struct {
	ArticleId               string    `json:"articleId"`
}
//空のオブジェクト
//object := Object{}

//記事一覧取得
func GetArticlesSql(userId, page, perPage string) *ArticlesAndPagination {
	userIdInt, userIdErr := strconv.Atoi(userId)
	if userIdErr != nil {
		userIdInt = 0
	}
	pageInt, pageErr := strconv.Atoi(page)
	if pageErr != nil {
		pageInt = 1
	}
	perPageInt, perPageErr := strconv.Atoi(perPage)
	if perPageErr != nil {
		perPageInt = 20
	}

	offset := (pageInt * perPageInt) - perPageInt

	articles := []Article{}
	selectCmd :=
		"SELECT articles.id,title,users.name,COALESCE(avatar_id,''),articles.created_at,articles.updated_at "+
			"FROM articles " +
			"LEFT JOIN users " +
			"ON articles.user_id = users.user_id "
		if userIdInt != 0 {
			selectCmd += "WHERE users.user_id = $1 "
		} else {
			selectCmd += "WHERE users.user_id <> $1 "
		}
		selectCmd += "ORDER BY updated_at desc LIMIT $2 OFFSET $3;"

		stmt, _ := Db.Prepare(selectCmd)
		defer stmt.Close()
	
		rows, _ := stmt.Query(userIdInt, perPageInt, offset)
		defer rows.Close()
		for rows.Next() {
			var p Article
			err := rows.Scan(
				&p.ArticleId,
				&p.ArticleTitle,
				&p.UserName,
				&p.AvatarId,
				&p.CreatedAt,
				&p.UpdatedAt)

			if err != nil {
				log.Fatalln(err)
			}
			articles = append(articles, p)
		}

	//ページ取得の巻
	var articlesCount int
		selectCountCmd := "select count(*) from articles "
		selectCountErr := Db.QueryRow(selectCountCmd).Scan(&articlesCount)
		if selectCountErr != nil {
			log.Fatalln(selectCountErr)
		}
	//
		return &ArticlesAndPagination{articles, Pagination{pageInt,perPageInt,articlesCount}}
}

//記事作成
func PostArticleSql(accessToken,articleTitle,content string) *ArticleEdit {
	selectAccessTokenCmd :=
		"SELECT user_id " +
			"FROM logins " +
			"WHERE uuid = $1;"

	var userId string
	selectAccessTokenErr := Db.QueryRow(selectAccessTokenCmd, accessToken).Scan(&userId)
	if selectAccessTokenErr != nil {
		log.Fatalln(selectAccessTokenErr)
	}

	//記事作成
	insertCmd :=
		"INSERT INTO articles (id,user_id,title,content,created_at,updated_at) " +
			"VALUES ($1,$2,$3,$4,$5,$6) RETURNING id;"
	var newArticleId ArticleEdit
	id,_ := MakeRandomStr(20)
	insertErr := Db.QueryRow(insertCmd, id,userId,articleTitle,content,  time.Now(), time.Now()).Scan(
		&newArticleId.ArticleId)
	if insertErr != nil {
		log.Fatal(insertErr)
	}
	return &newArticleId
}

 //記事取得
 func GetArticleByIdSql(articleId string) *ArticleDetail {
	selectArticleById :=
	"SELECT articles.id,title,content,users.name,COALESCE(avatar_id,''),articles.created_at,articles.updated_at "+
			"FROM articles " +
			"LEFT JOIN users " +
			"ON articles.user_id = users.user_id "+
			"where articles.id = $1;"
		var p ArticleDetail
		selectArticleByIdErr := Db.QueryRow(selectArticleById,articleId).Scan(
			&p.ArticleId,
			&p.ArticleTitle,
			&p.Content,
			&p.UserName,
			&p.AvatarId,
			&p.CreatedAt,
			&p.UpdatedAt)
		if selectArticleByIdErr != nil {
			log.Fatalln(selectArticleByIdErr)
		}
		return &p
  }

//記事更新
func PutArticleSql(accessToken,articleId,articleTitle,content string) *ArticleEdit {
	updateArticle :=
	"UPDATE articles " +
		"SET title = $1,content = $2,updated_at = $3 " +//タイトルとコンテンツと更新日をアプデ
		"WHERE id = $4;"
		_,_ = Db.Exec(updateArticle, articleTitle,content,time.Now(),articleId)
	return &ArticleEdit{articleId}
}

//記事削除
func DeleteArticleSql(accessToken,articleId string) {
//なんも返さない
	cmd := `
		delete
		from articles
		where id = $1
		`
	_,_ = Db.Exec(cmd, articleId)
}

//ランダム文字列作成
func MakeRandomStr(digit uint32) (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// 乱数を生成
	b := make([]byte, digit)
	if _, err := rand.Read(b); err != nil {
			return "", errors.New("unexpected error...")
	}

	// letters からランダムに取り出して文字列を生成
	var result string
	for _, v := range b {
			// index が letters の長さに収まるように調整
			result += string(letters[int(v)%len(letters)])
	}
	return result, nil
}