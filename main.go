package main

import (
	"fmt"

	"github.com/GrowthOdyssey/TechBoard-BE/config"
)

func main() {
	// config確認
	fmt.Println(config.Config.Port)
	fmt.Println(config.Config.SqlDriver)
	fmt.Println(config.Config.DbName)
}