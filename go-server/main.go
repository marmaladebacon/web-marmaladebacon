package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//e.Static("/", "../vue-client/dist")
	//e.Static("/wp", "../vue-client/web-projects")
	e.Use(middleware.Static("../vue-client/dist"))
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:  "../vue-client/web-projects",
		HTML5: true,
	}))

	e.GET("/", func(c echo.Context) error {
		return c.File("../vue-client/dist/index.html")
	})
	e.Logger.Fatal(e.Start(":8000"))
}
