package mw

import (
	"github.com/labstack/echo/v4"
	"strings"
)

func Ignore(extensions ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			for _, ext := range extensions {
				if strings.HasSuffix(c.Request().URL.Path, ext) {
					return echo.NotFoundHandler(c)
				}
			}
			return next(c)
		}
	}
}
