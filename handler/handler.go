package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

// MainPage の説明
func MainPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	}
}
