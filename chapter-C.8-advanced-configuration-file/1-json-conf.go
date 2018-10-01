package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
	"net/http"
)

func main() {
	e := echo.New()

	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.SetConfigName("app.config")

	err := viper.ReadInConfig()
	if err != nil {
		e.Logger.Fatal(err)
	}

	fmt.Println("Starting", viper.GetString("appName"))

	e.GET("/index", func(c echo.Context) (err error) {
		return c.JSON(http.StatusOK, true)
	})

	e.Logger.Fatal(e.Start(":" + viper.GetString("server.port")))
}
