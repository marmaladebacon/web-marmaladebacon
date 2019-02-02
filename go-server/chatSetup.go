package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// serveWsChat handles websocket requests from peer
func serveWsChat(chatHub *ChatHub, w http.ResponseWriter, r *http.Request) {
	// upgrading a HTTP connection to a WebSocket connection.
	/*
	 1. client sends a HTTP request requesting the server upgrade the connection used for the HTTP request to the WebSocket protocol.
	 2. The server inspects the request and if no errors are encountered, the server sends an HTTP response agreeing to upgrade the connection.
	 3. Going forward, client and server now use the WebSocket protocol over the network connection.
	*/
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// We create instances of the ChatClient type to manage the connection.
	client := &ChatClient{
		room:    "lobby",
		chatHub: chatHub,
		conn:    conn,
		send:    make(chan []byte, 1024),
	}

	//We send the instance of the client to the registration channel for chatHub
	client.chatHub.register <- client

	// Allow collection of memory referenced by the caller by doing all work in new goroutines
	// Start the client write and read goroutines to start read and write work for messages
	go client.writePump()
	go client.readPump()
	//go client.writeClientMsgPump()

	cWelcome := ToClientMsg{Cat: "admin", Text: "Welcome to marmaladebacon's chat hub!"}
	bytes, berr := json.Marshal(cWelcome)
	if berr != nil {
		panic(berr)
	}
	client.send <- bytes

	time.AfterFunc(1200*time.Millisecond, func() {
		sendToClientMsg(client, "admin", "type /help for a list of commands")
	})
}
