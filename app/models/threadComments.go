package models

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"
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

//スレッド個別取得時
type ThreadComment struct {
	Id        string    `json:"commentId"`
	Text      string    `json:"commentTitle"`
	UserId    string    `json:"userId"`
	UserName  string    `json:"userName"`
	SessionId string    `json:"sessionId"`
	CreatedAt time.Time `json:"createdAt"`
}

//コメント作成時
type CommentAndUser struct {
	Id        string    `json:"commentId"`
	Text      string    `json:"commentTitle"`
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

//スレッド個別取得時コメント取得
func GetCommentsByThreadIdSql(threadId string, Db *sql.DB) (*[]ThreadComment, int) {

	//コメントとユーザー取得
	stmt, _ := Db.Prepare(
		"SELECT id, text, COALESCE(thread_comments.user_id,''), COALESCE(users.name,''), COALESCE(session_id,''), thread_comments.created_at " +
			"FROM thread_comments LEFT JOIN users ON thread_comments.user_id = users.user_id " +
			"WHERE thread_id = $1;")
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

//コメント登録
func PostCommentsSql(threadId, userId, sessionId, commentTitle string) *CommentAndUser {
	threadIdInt, threadIdErr := strconv.Atoi(threadId)
	if threadIdErr != nil {
		log.Fatal(threadIdErr)
	}
	//コメント登録
	insertCmd :=
		"INSERT INTO thread_comments (user_id,thread_id,session_id,text,created_at,updated_at) " +
			"VALUES ($1,$2,$3,$4,$5,$6) RETURNING *;"
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

	updateThread :=
		"UPDATE threads " +
			"SET updated_at = $1 " +
			"WHERE id = $2;"
	upd, updateThreadErr := Db.Prepare(updateThread)
	if updateThreadErr != nil {
		log.Fatalln(updateThreadErr)
	}
	upd.Exec(newComment.UpdatedAt, threadIdInt)

	//作成したコメントとユーザー名取得
	selectUserName :=
		"SELECT thread_comments.id, text, " +
			"COALESCE(users.name,''), COALESCE(thread_comments.session_id,''), thread_comments.created_at " +
			"FROM thread_comments LEFT JOIN users ON thread_comments.user_id = users.user_id " +
			"where thread_comments.id = $1;"
	var commentAndUser CommentAndUser
	selectUserNameErr := Db.QueryRow(selectUserName, newComment.Id).Scan(
		&commentAndUser.Id,
		&commentAndUser.Text,
		&commentAndUser.UserName,
		&commentAndUser.SessionId,
		&commentAndUser.CreatedAt,
	)
	if selectUserNameErr != nil {
		log.Fatalln(selectUserNameErr)
	}
	return &commentAndUser
}
