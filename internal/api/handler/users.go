package handler

import (
	"net/http"
	"waitly/internal/api"
	"waitly/internal/api/data"

	"github.com/labstack/echo/v4"
)

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

func GetUsers(c echo.Context) error {
	db, err := api.GetDB()

	if err != nil {
		return err
	}

	query := data.GetAllUsersQuery()
	rows, err := db.Query(query)
	if err != nil {
		return err
	}

	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "failed to scan user: "+err.Error())
		}
		users = append(users, user)
	}

	return c.JSON(http.StatusOK, users)
}
