package lib

import (
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/echoview-v4"
	"html/template"
	"reflect"
	"strings"
	"time"
)

func Renderer() *echoview.ViewEngine {
	return echoview.New(goview.Config{
		Root:         "views",
		Extension:    ".html",
		Master:       "layouts/master",
		Funcs:        template.FuncMap{
			// splits a string
			"split": func(val, sep string) []string {
				return strings.Split(val, sep)
			},
			// joins a string
			"join": func(a, b, join string) string {
				return a + join + b
			},
			// negation
			"not": func(val bool) bool {
				return !val
			},
			// time format
			"time": func(val time.Time, layout string) string {
				return val.Format(layout)
			},
			// string larger than
			"larger_than": func(s string, n int) bool {
				return len(s) > n
			},
			// string limiter
			"lim": func(s string, n int) string {
				if len(s) > n {
					return s[:n] + "..."
				}
				return s
			},
			// length
			//"len": func(v interface{}) int {
			//	val := reflect.ValueOf(v)
			//	if val.Kind() == reflect.Ptr {
			//		val = val.Elem()
			//	}
			//	switch val.Kind() {
			//	case reflect.String:
			//	case reflect.Slice:
			//	case reflect.Array:
			//		return val.Len()
			//	}
			//	return -1
			//},
			// check if array is empty
			"empty": func(arr interface{}) bool {
				val := reflect.ValueOf(arr)
				if val.Kind() == reflect.Ptr {
					val = val.Elem()
				}
				if val.Kind() == reflect.Array || val.Kind() == reflect.Slice {
					return val.Len() == 0
				}
				return true
			},
		},
		DisableCache: true,
	})
}