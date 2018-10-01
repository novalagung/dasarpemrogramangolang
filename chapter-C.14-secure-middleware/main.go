package main

import (
	"github.com/labstack/echo"
	"github.com/unrolled/secure"
	"net/http"
)

func main() {
	e := echo.New()

	secureMiddleware := secure.New(secure.Options{
		AllowedHosts:            []string{"localhost:9000", "www.google.com"},
		FrameDeny:               true,
		CustomFrameOptionsValue: "SAMEORIGIN",
		ContentTypeNosniff:      true,
		BrowserXssFilter:        true,
	})

	e.Use(echo.WrapMiddleware(secureMiddleware.Handler))

	e.GET("/index", func(c echo.Context) error {
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")

		return c.String(http.StatusOK, "Hello")
	})

	e.Logger.Fatal(e.StartTLS(":9000", "server.crt", "server.key"))
}
