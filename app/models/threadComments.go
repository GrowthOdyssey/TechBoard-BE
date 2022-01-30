package models

import (
	"fmt"
	"time"
)

// スレッドコメント
type ThreadComment struct {
	Id        int       `json:"id"`
	UserId    string    `json:"userId"`
	ThradId   int       `json:"threadId"`
	SessionId string    `json:"sessionId"`
	Text      string    `json:"title"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (c *ThreadComment) ThreadCommentReceiver() {
	fmt.Println(c.Id, c.Text)
}

func ThreadCommentFunction(id string) {
	fmt.Println(id)
}
