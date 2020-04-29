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

func BrowseDir(root string, ignore []string) echo.MiddlewareFunc {
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

			var directories, files []echo.Map
			for _, item := range items {
				if strings.HasPrefix(item.Name(), ".") {
					continue
				}

				cont := false
				for _, ext := range ignore {
					if strings.HasSuffix(item.Name(), ext) {
						cont = true
						break
					}
				}
				if cont {
					continue
				}

				if item.IsDir() {
					directories = append(directories, echo.Map{
						"name": item.Name(),
						"size": item.Size(),
						"mod": item.ModTime(),
					})
				} else {
					files = append(files, echo.Map{
						"name": item.Name(),
						"size": item.Size(),
						"mod": item.ModTime(),
					})
				}
			}

			dirGrid := false
			if _, err := c.Cookie("cdragon_dir_grid"); err == nil {
				dirGrid = true
			}

			fileList := false
			if _, err := c.Cookie("cdragon_file_list"); err == nil {
				fileList = true
			}


			return c.Render(http.StatusOK, "browse", echo.Map{
				"directories": directories,
				"files": files,
				"current": echo.Map{
					"path": p,
					"settings": echo.Map{
						"dir_grid": dirGrid,
						"file_list": fileList,
					},
				},
			})
		}
	}
}
