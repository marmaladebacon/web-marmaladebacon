package main

import (
	"fmt"
	"time"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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

	// Chat App Setup
	chatHub := makeHub()
	go chatHub.run()
	e.GET("/ws/chat", func(c echo.Context) error {
		// http.ResponseWriter, r *http.Request
		r := c.Request
		w := &c.Response.Writer
		serveWsChat(chatHub, w, r)
	})

	// End

	e.Logger.Fatal(e.Start(":8080"))

	//setup perFrame duration 30fps so 1 second / 30
	perFrameDuration := time.Duration(time.Second / 30)
	/*
		'NewTicker returns a new Ticker containing a channel that will send the time with a period specified by the duration argument. It adjusts the intervals or drops ticks to make up for slow receivers. The duration d must be greater than zero; if not, NewTicker will panic. Stop the ticker to release associated resources.'

		We're setting up our game loop here
	*/
	clk := time.NewTicker(perFrameDuration)

	for {
		select {
		case <-clk.C:
			//loop
		}
	}
}
