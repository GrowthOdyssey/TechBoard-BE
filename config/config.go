package config

import (
	"fmt"

	"gopkg.in/go-ini/ini.v1"
)

// config構造体
type ConfigList struct {
	Port string
	SqlDriver string
	DbName string
}

var Config ConfigList

// 初期化関数
// main関数より前に呼ばれる
func init() {
	LoadConfig()
}

// config.iniを読み込む
func LoadConfig() {
	config, err := ini.Load("config.ini")

	if err != nil {
		fmt.Println(err)
	}

	Config = ConfigList{
		Port:      config.Section("server").Key("port").MustString("8080"),
		SqlDriver: config.Section("db").Key("driver").String(),
		DbName:    config.Section("db").Key("name").String(),
	}
}