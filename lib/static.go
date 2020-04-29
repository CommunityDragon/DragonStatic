package lib

import (
	"github.com/GeertJohan/go.rice"
	"github.com/labstack/echo/v4"
	"mime"
	"net/http"
	"path/filepath"
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

		var mimeType string
		if ext := filepath.Ext(name); ext != "" {
			mimeType = mime.TypeByExtension(ext)
		}
		if mimeType == "" {
			http.DetectContentType(box.MustBytes(name))
		}

		return c.Stream(200, mimeType, file)
	}
}
