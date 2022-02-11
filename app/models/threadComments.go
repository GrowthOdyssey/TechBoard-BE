package models

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
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
	Id        string    `json:"commentId"`
	Text      string    `json:"commentTitle"`
	ThreadId  string    `json:"threadId"`
	UserId    string    `json:"userId"`
	UserName  string    `json:"userName"`
	SessionId string    `json:"sessionId"`
	CreatedAt time.Time `json:"createdAt"`
}

func (c *ThreadComment) ThreadCommentReceiver() {
	fmt.Println(c.Id, c.Text)
}

func ThreadCommentFunction(id string) {
	fmt.Println(id)
}

func GetCommentsByThreadIdSql(threadId string, Db *sql.DB) (*[]ThreadComment, int) {

	stmt, _ := Db.Prepare("SELECT id, text, COALESCE(thread_comments.user_id,''), COALESCE(users.name,''), COALESCE(session_id,''), thread_comments.created_at " +
		"FROM thread_comments LEFT JOIN users ON thread_comments.user_id = users.user_id WHERE thread_id = $1;")
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

func PostCommentsSql(threadId, userId, sessionId, commentTitle string) *CommentAndThreadAndUser {
	connection := "user=test_user dbname=" + config.Config.DbName + " password=password sslmode=disable"
	Db, _ = sql.Open(config.Config.SqlDriver, connection)
	defer Db.Close()
	threadIdInt, threadIdErr := strconv.Atoi(threadId)
	if threadIdErr != nil {
		log.Fatal(threadIdErr)
	}
	insertCmd := "INSERT INTO thread_comments (user_id,thread_id,session_id,text,created_at,updated_at) VALUES ($1,$2,$3,$4,$5,$6) RETURNING *;"
	var newComment Comment
	insertErr := Db.QueryRow(insertCmd, userId, threadIdInt, sessionId, commentTitle, time.Now(), time.Now()).Scan(
		&newComment.Id,
		&newComment.UserId,
		&newComment.ThreadId,
		&newComment.SessionId,
		&newComment.Text,
		&newComment.CreatedAt,
		&newComment.UpdatedAt)
	if insertErr != nil {
		log.Fatal(insertErr)
	}

	selectUserName := "SELECT thread_comments.id, text, thread_id, COALESCE(thread_comments.user_id,''), " +
		"COALESCE(users.name,''), COALESCE(thread_comments.session_id,''), thread_comments.created_at " +
		"FROM thread_comments LEFT JOIN users ON thread_comments.user_id = users.user_id " +
		"where thread_comments.id = $1;"
	var commentAndThreadAndUser CommentAndThreadAndUser
	selectUserNameErr := Db.QueryRow(selectUserName, newComment.Id).Scan(
		&commentAndThreadAndUser.Id,
		&commentAndThreadAndUser.Text,
		&commentAndThreadAndUser.ThreadId,
		&commentAndThreadAndUser.UserId,
		&commentAndThreadAndUser.UserName,
		&commentAndThreadAndUser.SessionId,
		&commentAndThreadAndUser.CreatedAt,
	)
	if selectUserNameErr != nil {
		log.Fatalln(selectUserNameErr)
	}
	return &commentAndThreadAndUser
}
