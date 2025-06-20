package main

import (
	"fmt"
	"log"
	"waitly/internal/api/handler"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("Environment loaded!")

	e := echo.New()

	e.GET("/users", handler.GetUsers)

	e.Logger.Fatal(e.Start(":8080"))

}
