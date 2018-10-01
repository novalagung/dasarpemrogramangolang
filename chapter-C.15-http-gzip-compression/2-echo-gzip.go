package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"io"
	"os"
)

func main() {
	e := echo.New()

	e.Use(middleware.Gzip())

	e.GET("/image", func(c echo.Context) error {
		f, err := os.Open("sample.png")
		if err != nil {
			return err
		}

		_, err = io.Copy(c.Response(), f)
		if err != nil {
			return err
		}

		return nil
	})

	e.Logger.Fatal(e.Start(":9000"))
}
