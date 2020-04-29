package mw

import (
	"github.com/labstack/echo/v4"
	"strings"
)

func NoHidden(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if strings.Contains(c.Request().URL.Path, "/.") {
			return echo.NotFoundHandler(c)
		}
		return next(c)
	}
}