package main

import (
	"net/http"

	"github.com/dookie1983/myGoTest/fizzBuzz"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/fizzbuzz/:id", func(c echo.Context) error {
		id := c.Param("id")
		return c.String(http.StatusOK, fizzbuzz.FizzBuzz(id))
	})
	e.Logger.Fatal(e.Start(":1323"))
}
