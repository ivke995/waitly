package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type CreateDBRequest struct {
	Name string `json:"name"`
}

func CreateUserDatabase(c echo.Context) error {
	var req CreateDBRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}
	if req.Name == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Database name is required")
	}

	orgSlug := os.Getenv("TURSO_ORG_NAME")
	group := os.Getenv("TURSO_GROUP_NAME")
	token := os.Getenv("TURSO_ORG_TOKEN")

	fmt.Print(orgSlug, group, token)

	if orgSlug == "" || group == "" || token == "" {
		return echo.NewHTTPError(http.StatusInternalServerError, "Missing required env vars")
	}

	bodyData := map[string]string{
		"name":  req.Name,
		"group": group,
	}
	bodyJson, err := json.Marshal(bodyData)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to encode JSON")
	}

	url := fmt.Sprintf("https://api.turso.tech/v1/organizations/%s/databases", orgSlug)
	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyJson))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create request")
	}
	httpReq.Header.Set("Authorization", "Bearer "+token)
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, "Turso API request failed: "+err.Error())
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		return echo.NewHTTPError(resp.StatusCode, string(respBody))
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "Database created successfully",
		"turso":   json.RawMessage(respBody),
	})
}
