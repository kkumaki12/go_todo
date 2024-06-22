package main

import (
	"fmt"
	"go_todo/app/controllers"
	"go_todo/app/models"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()
}
