# Summary of Chat App
1. Front End App
1. Back End App  
  * Chat Setup: 
    * (Outside of ChatSetup) Initialize ChatHub in main.go and run as a go routine
    *
  * Chat Hub: Directs message traffic via broadcast channel to clients
  * Chat Client


## Chat Setup
* Set up a route on main
```go
e.GET("/ws/chat", func(c echo.Context) error {
  // http.ResponseWriter, r *http.Request
  r := c.Request()
  w := c.Response()
  serveWsChat(chatHub, w, r)
  return nil
})
```
* Upgrade an incoming connection in ChatSetup
```go
//w http.ResponseWriter, r *http.Request
conn, err := upgrader.Upgrade(w, r, nil)
```

* Create an instance of chatClient, complete with a reference to the existing ChatHub
```go
// We create instances of the ChatClient type to manage the connection.
client := &ChatClient{
  chatHub: chatHub,
  conn:    conn,
  send:    make(chan []byte, 256),
}
```

* Register the chat client against the chatHub, and run the message directors
```go
//We send the instance of the client to the registration channel for chatHub
client.chatHub.register <- client

// Allow collection of memory referenced by the caller by doing all work in new goroutines
// Start the client write and read goroutines to start read and write work for messages
go client.writePump()
go client.readPump()
```