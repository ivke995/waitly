package main

import (
	"fmt"
	"waitly/internal/api/handler"

	"github.com/labstack/echo/v4"
)

func main() {

	fmt.Println("Environment loaded!")

	e := echo.New()

	e.GET("/users", handler.GetUsers)
	e.POST("/user-db-create", handler.CreateUserDatabase)
	e.POST("/signup", handler.InsertUser)

	e.Logger.Fatal(e.Start(":8080"))

}
