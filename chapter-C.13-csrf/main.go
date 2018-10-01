package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"html/template"
	"net/http"
)

type M map[string]interface{}

const CSRF_TOKEN_HEADER = "X-Csrf-Token"
const CSRF_KEY = "csrf_token"

func main() {
	tmpl := template.Must(template.ParseGlob("./*.html"))

	e := echo.New()

	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "header:" + CSRF_TOKEN_HEADER,
		ContextKey:  CSRF_KEY,
	}))

	e.GET("/index", func(c echo.Context) error {
		data := make(M)
		data[CSRF_KEY] = c.Get(CSRF_KEY)
		return tmpl.Execute(c.Response(), data)
	})

	e.POST("/sayhello", func(c echo.Context) error {
		data := make(M)
		if err := c.Bind(&data); err != nil {
			return err
		}

		message := fmt.Sprintf("hello %s", data["name"])
		return c.JSON(http.StatusOK, message)
	})

	e.Logger.Fatal(e.Start(":9000"))
}
