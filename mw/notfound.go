package mw

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func NotFound(c echo.Context) error {
	return c.Render(http.StatusNotFound, "404", nil)
}
