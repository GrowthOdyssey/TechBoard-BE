package models

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/GrowthOdyssey/TechBoard-BE/config"
)

// スレッドコメント
type Comment struct {
	Id        string    `json:"id"`
	UserId    string    `json:"userId"`
	ThreadId  string    `json:"threadId"`
	SessionId string    `json:"sessionId"`
	Text      string    `json:"title"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type ThreadComment struct {
	Id        string    `json:"commentId"`
	Text      string    `json:"commentTitle"`
	UserId    string    `json:"userId"`
	UserName  string    `json:"userName"`
	SessionId string    `json:"sessionId"`
	CreatedAt time.Time `json:"createdAt"`
}

type CommentAndThreadAndUser struct {
	Id          string    `json:"commentId"`
	Text        string    `json:"commentTitle"`
	ThreadId    string    `json:"threadId"`
	ThreadTitle string    `json:"threadTitle"`
	UserId      string    `json:"userId"`
	UserName    string    `json:"userName"`
	SessionId   string    `json:"sessionId"`
	CreatedAt   time.Time `json:"createdAt"`
}

func (c *ThreadComment) ThreadCommentReceiver() {
	fmt.Println(c.Id, c.Text)
}

func ThreadCommentFunction(id string) {
	fmt.Println(id)
}

func GetCommentsByThreadIdSql(threadId string, Db *sql.DB) (*[]ThreadComment, int) {

	stmt, _ := Db.Prepare("select id, text, COALESCE(thread_comments.user_id,''), COALESCE(users.name,''), COALESCE(session_id,''), thread_comments.created_at from thread_comments LEFT JOIN users ON thread_comments.user_id = users.user_id where thread_id = $1;")
	defer stmt.Close()
	var comments []ThreadComment
	rows, err := stmt.Query(threadId)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()
	for rows.Next() {
		var p ThreadComment
		err := rows.Scan(
			&p.Id,
			&p.Text,
			&p.UserId,
			&p.UserName,
			&p.SessionId,
			&p.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		comments = append(comments, p)
	}
	return &comments, len(comments)
}

