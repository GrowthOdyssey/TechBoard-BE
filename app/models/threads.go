package models

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

type ModelErrMsg struct {
	ErrorMessage string `json:"message"`
}

// スレッド
type Thread struct {
	Id               string    `json:"threadId"`
	UserId           string    `json:"userId"`
	ThreadCategoryId string    `json:"categoryId"`
	Title            string    `json:"threadTitle"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
	FirstComment     string    `json:"firstComment"`
	CommentsCount    int       `json:"commentsCount"`
}

type Pagination struct {
	Page    int `json:"page"`
	PerPage int `json:"perPage"`
	Total   int `json:"total"`
}

//スレッド一覧取得用
type ThreadsAndPagination struct {
	Threads    []Thread   `json:"threads"`
	Pagination Pagination `json:"pagination"`
}

//スレッド作成用
type NewThreadId struct {
	ThreadId string `json:"threadId"`
}

//スレッド個別取得用
type ThreadAndComments struct {
	Id               string           `json:"threadId"`
	UserId           string           `json:"userId"`
	ThreadCategoryId string           `json:"categoryId"`
	Title            string           `json:"threadTitle"`
	CreatedAt        time.Time        `json:"createdAt"`
	UpdatedAt        time.Time        `json:"updatedAt"`
	Comments         *[]ThreadComment `json:"comments"`
	CommentsCount    int              `json:"commentsCount"`
}

//スレッドカテゴリー一覧取得用
type ThreadsCategory struct {
	Id           string `json:"categoryId"`
	CategoryName string `json:"categoryName"`
}
type ThreadsCategories struct {
	ThreadsCategories *[]ThreadsCategory `json:"categories"`
}

func (t *Thread) ThreadReceiver() {
	fmt.Println(t.Id, t.Title)
}

func SampleFunction(id string) {
	fmt.Println(id)
}

//スレッド一覧取得
func GetThreadsSql(categoryId, page, perPage string) *ThreadsAndPagination {
	categoryIdInt, categoryIdErr := strconv.Atoi(categoryId)
	if categoryIdErr != nil {
		categoryIdInt = 0
	}
	pageInt, pageErr := strconv.Atoi(page)
	if pageErr != nil {
		pageInt = 1
	}
	perPageInt, perPageErr := strconv.Atoi(perPage)
	if perPageErr != nil {
		perPageInt = 20
	}
	// DB接続

	// スレッド一覧取得 一番古いコメントを結合
	var threads []Thread
	selectCmd :=
		"SELECT threads.*, COALESCE(oldest_comment.text,''), COALESCE(comments_count.count,0) " +
			"FROM threads " +
			"LEFT JOIN " +
			"(SELECT * " +
			"FROM thread_comments " +
			"WHERE NOT EXISTS " +
			"(SELECT 1 " +
			"FROM thread_comments sub " +
			"WHERE thread_comments.thread_id = sub.thread_id " +
			"AND thread_comments.created_at > sub.created_at)) " +
			"oldest_comment " +
			"ON threads.id = oldest_comment.thread_id " +
			"LEFT JOIN " +
			"(SELECT count(*), thread_id " +
			"FROM thread_comments " +
			"GROUP BY thread_id) comments_count " +
			"ON threads.id = comments_count.thread_id "
	if categoryIdInt != 0 {
		selectCmd += "WHERE thread_category_id = $1 "
	} else {
		selectCmd += "WHERE thread_category_id <> $1 "
	}
	selectCmd += "ORDER BY updated_at desc;"

	stmt, _ := Db.Prepare(selectCmd)
	defer stmt.Close()

	rows, _ := stmt.Query(categoryIdInt)
	defer rows.Close()
	for rows.Next() {
		var p Thread
		err := rows.Scan(
			&p.Id,
			&p.UserId,
			&p.ThreadCategoryId,
			&p.Title,
			&p.CreatedAt,
			&p.UpdatedAt,
			&p.FirstComment,
			&p.CommentsCount)

		if err != nil {
			log.Fatalln(err)
		}
		threads = append(threads, p)
	}

	return &ThreadsAndPagination{threads, Pagination{pageInt, perPageInt, len(threads)}}
}

//スレッド作成
func PostThreadSql(accessToken, threadTitle, categoryId string) *NewThreadId {
	selectAccessTokenCmd :=
		"SELECT user_id " +
			"FROM logins " +
			"WHERE uuid = $1;"

	var userId string
	selectAccessTokenErr := Db.QueryRow(selectAccessTokenCmd, accessToken).Scan(&userId)
	if selectAccessTokenErr != nil {
		log.Fatalln(selectAccessTokenErr)
	}

	//スレッド登録
	insertCmd :=
		"INSERT INTO threads (user_id,thread_category_id,title,created_at,updated_at) " +
			"VALUES ($1,$2,$3,$4,$5) RETURNING id;"
	categoryIdInt, categoryIdErr := strconv.Atoi(categoryId)
	if categoryIdErr != nil {
		categoryIdInt = 1
	}
	var newThreadId NewThreadId
	insertErr := Db.QueryRow(insertCmd, userId, categoryIdInt, threadTitle, time.Now(), time.Now()).Scan(
		&newThreadId.ThreadId)
	if insertErr != nil {
		log.Fatal(insertErr)
	}
	return &newThreadId
}

//スレッド個別取得
func GetThreadByIdSql(id string) *ThreadAndComments {

	//スレッド取得
	selectThreadById :=
		"SELECT * " +
			"from threads " +
			"where id = $1;"
	var thread Thread
	selectThreadByIdErr := Db.QueryRow(selectThreadById, id).Scan(
		&thread.Id,
		&thread.UserId,
		&thread.ThreadCategoryId,
		&thread.Title,
		&thread.CreatedAt,
		&thread.UpdatedAt)
	if selectThreadByIdErr != nil {
		log.Fatalln(selectThreadByIdErr)
	}

	threadComments, commentsCount := GetCommentsByThreadIdSql(id, Db)

	return &ThreadAndComments{
		thread.Id,
		thread.UserId,
		thread.ThreadCategoryId,
		thread.Title,
		thread.CreatedAt,
		thread.UpdatedAt,
		threadComments,
		commentsCount}
}

//スレッドカテゴリー一覧取得
func GetThreadsCategoriesSql() *ThreadsCategories {
	selectCmd := "SELECT id, name FROM thread_categories;"
	var threadsCategories []ThreadsCategory
	rows, err := Db.Query(selectCmd)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()
	for rows.Next() {
		var p ThreadsCategory
		err := rows.Scan(
			&p.Id,
			&p.CategoryName)
		if err != nil {
			log.Fatalln(err)
		}
		threadsCategories = append(threadsCategories, p)
	}
	return &ThreadsCategories{&threadsCategories}
}
