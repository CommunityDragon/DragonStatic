package main

//go:generate rice embed-go

import (
	"dragonstatic/lib"
	"dragonstatic/mw"
	"flag"
	"github.com/GeertJohan/go.rice"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"os"
	"strings"
)

var dir = ""
var address = ""
var ignore []string

func loadArgs() {
	addrPtr := flag.String("address", ":5050", "address to run it on")
	ignorePtr := flag.String("ignore", ".exe,.dll", "ignore extensions from being exposed")
	flag.Parse()
	dirArg := flag.Arg(0)

	if dirArg == "" {
		log.Fatal("error: static path not set")
	}
	if _, err := os.Stat(dirArg); os.IsNotExist(err) {
		log.Fatal(err.Error())
	}

	address = *addrPtr
	ignore = strings.Split(*ignorePtr, ",")
	dir = dirArg
}

func main() {
	loadArgs()
	renderer := lib.Renderer()
	renderer.SetFileHandler(lib.FileHandler(rice.MustFindBox("views")))
	echo.NotFoundHandler = mw.NotFound

	e := echo.New()
	e.Renderer = renderer

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Pre(middleware.RemoveTrailingSlash())
	e.GET("/*", lib.Static(rice.MustFindBox("static")))

	e.Use(mw.NoHidden)
	e.Use(mw.Ignore(ignore...))
	e.Use(mw.BrowseDir(dir, ignore))
	e.Use(middleware.Static(dir))

	e.Logger.Fatal(e.Start(address))
}
