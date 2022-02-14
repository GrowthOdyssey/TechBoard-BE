package models

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// モデル
// 構造体定義、DBにアクセスする関数、レシーバメソッドを記載する。

// struct

// ユーザー登録リクエスト構造体
type UserSignUpReq struct {
	UserId	string	`json:"userId"`
	Password	string	`json:"password"`
	UserName	string	`json:"userName"`
	AvatarId	string	`json:"avatarId"`
}

// ユーザーレスポンス構造体
type UserRes struct {
	UserId	string	`json:"userId"`
	UserName	string	`json:"userName"`
	AvatarId	string	`json:"avatarId"`
	AccessToken	string	`json:"accessToken"`
	CreatedAt	string	`json:"createdAt"`
}

// ユーザーログインリクエスト構造体
type UserLoginReq struct {
	UserId	string	`json:"userId"`
	Password	string	`json:"password"`
}

// public function

// ユーザーを登録する
func (u *UserSignUpReq) RegisterUser() (userRes UserRes, err error) {
	cmd := `
		insert into users (user_id, name, password, avatar_id, created_at, updated_at)
		values ($1, $2, $3, $4, $5, $6)
	`

	_, err = Db.Exec(
		cmd, u.UserId, u.UserName, encrypt(u.Password), u.AvatarId, time.Now(), time.Now())
	if err != nil {
		fmt.Println(err)
	}

	// 登録した最新レコードを取得する
	cmd = `
		select user_id, name, avatar_id, created_at
		from users
		order by created_at desc
		limit 1
	`

	err = Db.QueryRow(cmd).Scan(
		&userRes.UserId,
		&userRes.UserName,
		&userRes.AvatarId,
		&userRes.CreatedAt,
	)

	return userRes, err
}

// ログインする
func Login(userId string) (accessToken string, err error) {
	cmd := `
		insert into logins (uuid, user_id, created_at)
		values ($1, $2, $3)
	`

	_, err = Db.Exec(cmd, createUuid(), userId, time.Now())
	if err != nil {
		fmt.Println(err)
	}

	// 登録した最新レコードを取得する
	cmd = `
		select uuid
		from logins
		order by created_at desc
		limit 1
	`

	err = Db.QueryRow(cmd).Scan(&accessToken)

	return accessToken, err
}

// ログインチェックをする
func (u *UserLoginReq) CheckLogin() (isOk bool, userRes UserRes, err error) {
	cmd := `
		select user_id, name, avatar_id, created_at
		from users
		where user_id = $1 and password = $2
	`

	err = Db.QueryRow(cmd, u.UserId, encrypt(u.Password)).Scan(
		&userRes.UserId,
		&userRes.UserName,
		&userRes.AvatarId,
		&userRes.CreatedAt,
	)

	if err != nil || userRes.UserId == "" {
		isOk = false
	} else {
		isOk = true
	}

	return isOk, userRes, err
}

// ログアウトする
func Logout(accessToken string) (err error) {
	cmd := `
		delete
		from logins
		where uuid = $1
	`

	_, err = Db.Exec(cmd, accessToken)

	return err
}

// private function

// 文字列を暗号化する
func encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}

// UUIDを生成する
func createUuid() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}