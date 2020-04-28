package lib

import (
	"fmt"
	"github.com/GeertJohan/go.rice"
	"github.com/foolin/goview"
	"path/filepath"
)

func FileHandler(riceBox *rice.Box) goview.FileHandler {
	box := riceBox.HTTPBox()

	return func (config goview.Config, tplFile string) (content string, err error) {
		path, err := filepath.Rel(".", tplFile + config.Extension)
		if err != nil {
			return "", fmt.Errorf("ViewEngine path:%v error: %v", path, err)
		}
		data, err := box.Bytes(path)
		if err != nil {
			return "", fmt.Errorf("ViewEngine render read name:%v, path:%v, error: %v", tplFile, path, err)
		}
		return string(data), nil
	}
}
