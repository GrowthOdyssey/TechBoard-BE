package main

import (
	"fmt"

	// init関数コールのため空importする
	_ "github.com/GrowthOdyssey/TechBoard-BE/app/models"
	_ "github.com/GrowthOdyssey/TechBoard-BE/config"

	"github.com/GrowthOdyssey/TechBoard-BE/app/controllers"
)

func main() {
	controllers.SetRouter()

	fmt.Println("APIサーバを起動します。")
	controllers.StartServer()
}