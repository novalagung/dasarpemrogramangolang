package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

type User struct {
	Name  string `json:"name" form:"name" query:"name"`
	Email string `json:"email" form:"email" query:"email"`
}

func main() {
	r := echo.New()

	r.Any("/user", func(c echo.Context) (err error) {
		u := new(User)
		if err = c.Bind(u); err != nil {
			return
		}

		return c.JSON(http.StatusOK, u)
	})

	// json payload
	// curl -X POST http://localhost:9000/user -H 'Content-Type: application/json' -d '{"name":"Joe","email":"joe@novalagung.com"}'

	// xml payload
	// curl -X POST http://localhost:9000/user -H 'Content-Type: application/xml' -d '<?xml version="1.0"?><Data><Name>Joe</Name><Email>joe@novalagung.com</Email></Data>'

	// form data
	// curl -X POST http://localhost:9000/user -d 'name=Joe' -d 'email=joe@novalagung.com'

	// query string
	// curl -X GET http://localhost:9000/user?name=Joe&email=joe@novalagung.com

	fmt.Println("server started at :9000")
	r.Start(":9000")
}
