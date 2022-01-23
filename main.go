package main

import (
	"fmt"

	// init関数コールのため空importする
	_ "github.com/GrowthOdyssey/TechBoard-BE/app/models"
	_ "github.com/GrowthOdyssey/TechBoard-BE/config"
)

func main() {
	fmt.Println("起動")
}