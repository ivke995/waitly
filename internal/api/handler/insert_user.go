package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"waitly/internal/api"
	"waitly/internal/api/data"

	"github.com/labstack/echo/v4"
)

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func sanitize(input string) string {
	s := strings.ToLower(input)
	s = strings.ReplaceAll(s, "@", "-")
	s = strings.ReplaceAll(s, ".", "-")
	re := regexp.MustCompile(`[^a-z0-9\-]+`)
	return re.ReplaceAllString(s, "")
}

func InsertUser(c echo.Context) error {
	db, err := api.GetDB()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get DB: "+err.Error())
	}

	var req CreateUserRequest
	if err := c.Bind(&req); err != nil || req.Username == "" || req.Email == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
	}

	result, err := db.Exec(data.InsertUserQuery(), req.Username, req.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to insert user: "+err.Error())
	}

	id, err := result.LastInsertId()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get user ID")
	}

	dbName := fmt.Sprintf("userdb-%d-%s-%s", id, sanitize(req.Username), sanitize(req.Email))

	payload, _ := json.Marshal(map[string]string{"name": dbName})
	resp, err := http.Post("http://localhost:8080/user-db-create", "application/json", bytes.NewBuffer(payload))
	if err != nil || (resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated) {
		return echo.NewHTTPError(http.StatusBadGateway, "Failed to create user DB")
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"id":       id,
		"username": req.Username,
		"email":    req.Email,
		"dbName":   dbName,
	})
}
