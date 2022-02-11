package models

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/GrowthOdyssey/TechBoard-BE/config"
	_ "github.com/lib/pq"
)

var Db *sql.DB
var err error

// 初期化関数
// main.goでimportするとmain関数より前に呼ばれる
func init() {
	// 本番モードはsslmode=requireにする
	connection := os.Getenv("DATABASE_URL")
	fmt.Println(connection, "yeah")
	if connection == "" {
		connection = "user=test_user dbname=" + config.Config.DbName + " password=password sslmode=disable"
	}
	Db, err = sql.Open(config.Config.SqlDriver, connection)

	if err != nil {
		fmt.Println(err)
	}

	// TODO aiharanaoya 以下、接続確認処理。後で消す
	// cmd := `
	// 	insert into users (name, email)
	// 	values ($1, $2)
	// `

	// _, err = Db.Exec(cmd, "name", "test@test.com")

	// if err != nil {
	// 	fmt.Println(err)
	// }
}
