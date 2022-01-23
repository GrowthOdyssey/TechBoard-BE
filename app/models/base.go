package models

import (
	"database/sql"
	"fmt"

	"github.com/GrowthOdyssey/TechBoard-BE/config"
	_ "github.com/lib/pq"
)

var Db *sql.DB
var err error

// 初期化関数
// main.goでimportするとmain関数より前に呼ばれる
func init() {
	// 本番モードはsslmode=requireにする
	connection := "user=test_user dbname=" + config.Config.DbName + " password=password sslmode=disable"
	Db, err = sql.Open(config.Config.SqlDriver, connection)

	if err != nil {
		fmt.Println(err)
	}
}