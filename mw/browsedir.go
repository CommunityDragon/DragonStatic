package mw

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func BrowseDir(root string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			p := c.Request().URL.Path
			if strings.HasSuffix(c.Path(), "*") {
				p = c.Param("*")
			}

			p, err = url.PathUnescape(p)
			if err != nil {
				return
			}

			name := filepath.Join(root, path.Clean("/"+p))

			file, err := os.Open(name)
			if err != nil {
				return next(c)
			}
			items, err := file.Readdir(-1)
			if err != nil {
				return next(c)
			}

			var directories, files []string
			for _, item := range items {
				if strings.HasPrefix(item.Name(), ".") {
					continue
				}

				if item.IsDir() {
					directories = append(directories, item.Name())
				} else {
					files = append(files, item.Name())
				}
			}

			return c.Render(http.StatusOK, "browse", echo.Map{
				"directories": directories,
				"files": files,
			})
		}
	}
}
