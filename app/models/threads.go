package models

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/GrowthOdyssey/TechBoard-BE/config"
)

// スレッド
type Thread struct {
	Id               string    `json:"threadId"`
	UserId           string    `json:"userId"`
	ThreadCategoryId string    `json:"categoryId"`
	Title            string    `json:"threadTitle"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}

type Pagination struct {
	Page    int `json:"page"`
	PerPage int `json:"perPage"`
	Total   int `json:"total"`
}

type ThreadsAndPagination struct {
	Threads    []Thread   `json:"threads"`
	Pagination Pagination `json:"pagination"`
}

type ThreadAndUser struct {
	Id               string    `json:"threadId"`
	UserId           string    `json:"userId"`
	ThreadCategoryId string    `json:"categoryId"`
	Title            string    `json:"threadTitle"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
	UserName         string    `json:"userName"`
}

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

func (t *Thread) ThreadReceiver() {
	fmt.Println(t.Id, t.Title)
}

func SampleFunction(id string) {
	fmt.Println(id)
}

func GetThreadsSql(page, perPage string) *ThreadsAndPagination {
	pageInt, pageErr := strconv.Atoi(page)
	if pageErr != nil {
		pageInt = 1
	}
	perPageInt, perPageErr := strconv.Atoi(perPage)
	if perPageErr != nil {
		perPageInt = 20
	}
	// DB接続
	connection := "user=test_user dbname=" + config.Config.DbName + " password=password sslmode=disable"
	Db, _ = sql.Open(config.Config.SqlDriver, connection)
	defer Db.Close()

	// スレッド一覧取得
	var threads []Thread
	selectCmd := "select * from threads;"
	rows, _ := Db.Query(selectCmd)
	defer rows.Close()
	for rows.Next() {
		var p Thread
		err := rows.Scan(
			&p.Id,
			&p.UserId,
			&p.ThreadCategoryId,
			&p.Title,
			&p.CreatedAt,
			&p.UpdatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		threads = append(threads, p)
	}

	/*
		// スレッド総件数取得
		selectCountCmd := "select count(*) from threads;"
		var count int
		err := Db.QueryRow(selectCountCmd).Scan(&count)
		if err != nil {
			log.Fatal(err)
		}
	*/

	return &ThreadsAndPagination{threads, Pagination{pageInt, perPageInt, len(threads)}}
}

func PostThreadSql(accessToken, threadTitle, categoryId string) *ThreadAndUser {
	connection := "user=test_user dbname=" + config.Config.DbName + " password=password sslmode=disable"
	Db, _ = sql.Open(config.Config.SqlDriver, connection)
	defer Db.Close()
	selectAccessTokenCmd := "select user_id from logins where uuid = $1;"
	var userId string
	selectAccessTokenErr := Db.QueryRow(selectAccessTokenCmd, accessToken).Scan(&userId)
	if selectAccessTokenErr != nil {
		log.Fatalln(selectAccessTokenErr)
	}

	insertCmd := "INSERT INTO threads (user_id,thread_category_id,title,created_at,updated_at) VALUES ($1,$2,$3,$4,$5) RETURNING *;"
	categoryIdInt, categoryIdErr := strconv.Atoi(categoryId)
	if categoryIdErr != nil {
		categoryIdInt = 1
	}
	var newThread Thread
	insertErr := Db.QueryRow(insertCmd, userId, categoryIdInt, threadTitle, time.Now(), time.Now()).Scan(
		&newThread.Id,
		&newThread.UserId,
		&newThread.ThreadCategoryId,
		&newThread.Title,
		&newThread.CreatedAt,
		&newThread.UpdatedAt)
	if insertErr != nil {
		log.Fatal(insertErr)
	}

	selectUserName := "select name from users where user_id = $1;"
	var userName string
	selectUserNameErr := Db.QueryRow(selectUserName, newThread.UserId).Scan(&userName)
	if selectUserNameErr != nil {
		log.Fatalln(selectUserNameErr)
	}

	return &ThreadAndUser{
		newThread.Id,
		newThread.UserId,
		newThread.ThreadCategoryId,
		newThread.Title,
		newThread.CreatedAt,
		newThread.UpdatedAt,
		userName}
}

func GetThreadByIdSql(id string) *ThreadAndComments {
	connection := "user=test_user dbname=" + config.Config.DbName + " password=password sslmode=disable"
	Db, _ = sql.Open(config.Config.SqlDriver, connection)
	defer Db.Close()
	selectThreadById := "select * from threads where id = $1;"
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

	/*
		selectCountThreadComments := "select count(*) from thread_comments where thread_id = $1;"
		var commentsCount int
		err := Db.QueryRow(selectCountThreadComments, id).Scan(&commentsCount)
		if err != nil {
			log.Fatal(err)
		}
	*/

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
