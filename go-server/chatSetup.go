package main

import (
	"log"
	"net/http"
)

// serveWsChat handles websocket requests from peer
func serveWsChat(chatHub *ChatHub, w http.ResponseWriter, r *http.Request) {
	// upgrading a HTTP connection to a WebSocket connection.
	/*
	 1. client sends a HTTP request requesting the server upgrade the connection used for the HTTP request to the WebSocket protocol.
	 2. The server inspects the request and if no errors are encountered, the server sends an HTTP response agreeing to upgrade the connection.
	 3. Going forward, client and server now use the WebSocket protocol over the network connection.
	*/
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// We create instances of the ChatClient type to manage the connection.
	client := &ChatClient{
		chatHub: chatHub,
		conn:    conn,
		send:    make(chan []byte, 256),
	}

	//We send the instance of the client to the registration channel for chatHub
	client.chatHub.register <- client

	// Allow collection of memory referenced by the caller by doing all work in new goroutines
	// Start the client write and read goroutines to start read and write work for messages
	go client.writePump()
	go client.readPump()
}
