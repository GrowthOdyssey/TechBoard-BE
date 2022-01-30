package models

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/GrowthOdyssey/TechBoard-BE/config"
)

// スレッド
type Thread struct {
	Id               int       `json:"threadId"`
	UserId           string    `json:"userId"`
	ThreadCategoryId int       `json:"categoryId"`
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

// type Sample struct {
// 	Test1 string `json:"test1"`
// 	Test2 int `json:"test2"`
// }

func (t *Thread) ThreadReceiver() {
	fmt.Println(t.Id, t.Title)
}

func SampleFunction(id string) {
	fmt.Println(id)
}

func GetThreadsSql(page, perPage int) *ThreadsAndPagination {
	connection := "user=test_user dbname=" + config.Config.DbName + " password=password sslmode=disable"
	Db, _ = sql.Open(config.Config.SqlDriver, connection)
	defer Db.Close()
	var threads []Thread
	selectCmd := "select * from threads"
	rows, _ := Db.Query(selectCmd)
	selectCountCmd := "select count(*) from threads"
	var count int
	err := Db.QueryRow(selectCountCmd).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
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

	return &ThreadsAndPagination{threads, Pagination{page, perPage, count}}
}
