package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/", "../vue-client/dist")

	e.GET("/", func(c echo.Context) error {
		return c.File("../vue-client/dist/index.html")
	})
	e.Logger.Fatal(e.Start(":8000"))
}
