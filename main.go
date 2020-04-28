package main

//go:generate rice embed-go

import (
	"dragonstatic/lib"
	"dragonstatic/mw"
	"flag"
	"github.com/GeertJohan/go.rice"
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"os"
)

var dir = ""

func loadArgs() {
	flag.Parse()
	dirArg := flag.Arg(0)

	if dirArg == "" {
		log.Fatal("error: static path not set")
	}
	if _, err := os.Stat(dirArg); os.IsNotExist(err) {
		log.Fatal(err.Error())
	}
	dir = dirArg
}

func main() {
	loadArgs()
	renderer := echoview.Default()
	renderer.SetFileHandler(lib.FileHandler(rice.MustFindBox("views")))
	echo.NotFoundHandler = mw.NotFound

	e := echo.New()
	e.Renderer = renderer

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/*", lib.Static(rice.MustFindBox("static")))

	e.Use(mw.NoHidden)
	e.Use(mw.BrowseDir(dir))
	e.Use(middleware.Static(dir))

	e.Logger.Fatal(e.Start(":3000"))
}
