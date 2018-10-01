package main

import (
	"github.com/labstack/echo"
	"net/http"
)

var ActionIndex = func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("from action index"))
}

var ActionHome = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("from action home"))
})

var ActionAbout = echo.WrapHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("from action about"))
}))

func main() {
	r := echo.New()

	r.GET("/index", echo.WrapHandler(http.HandlerFunc(ActionIndex)))
	r.GET("/home", echo.WrapHandler(ActionHome))
	r.GET("/about", ActionAbout)

	r.Start(":9000")
}
