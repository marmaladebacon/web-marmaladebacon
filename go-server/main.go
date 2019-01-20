package main

import (
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	upgrader = websocket.Upgrader{}
)

func hello(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		// Write
		err := ws.WriteMessage(websocket.TextMessage, []byte("Hello, Client!"))
		if err != nil {
			c.Logger().Error(err)
		}

		// Read
		_, msg, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Error(err)
		}
		fmt.Printf("%s\n", msg)
	}
}

func main() {
	e := echo.New()

	// enable when devel
	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Static("../vue-client/dist"))
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:  "../vue-client/web-projects",
		HTML5: true,
	}))

	e.GET("/", func(c echo.Context) error {
		return c.File("../vue-client/dist/index.html")
	})

	e.GET("/ws/hello", hello)

	e.Logger.Fatal(e.Start(":8080"))
}
