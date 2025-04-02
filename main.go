package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func main() {
	e := echo.New()
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}
	e.GET("/", func(c echo.Context) error {
		dbUrl := os.Getenv("TURSO_DATABASE_URL")
		authToken := os.Getenv("TURSO_AUTH_TOKEN")

		url := fmt.Sprintf("%s?authToken=%s", dbUrl, authToken)
		fmt.Print(url)

		db, err := sql.Open("libsql", url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to open db %s: %s", url, err)
			os.Exit(1)
		}
		defer db.Close()
		return c.String(http.StatusOK, "Hello,World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
