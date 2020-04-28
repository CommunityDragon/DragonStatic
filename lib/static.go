package lib

import (
	"github.com/GeertJohan/go.rice"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func Static(riceBox *rice.Box) echo.HandlerFunc {
	box := riceBox.HTTPBox()
	return func(c echo.Context) (err error) {
		name := strings.TrimPrefix(c.Request().RequestURI, "/")
		file, err := box.Open(name)
		if err != nil {
			return c.Render(http.StatusNotFound, "404", nil)
		}
		return c.Stream(200, http.DetectContentType(box.MustBytes(name)), file)
	}
}
