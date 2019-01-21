package main

import (
	"log"
	"net/http"
)

// serveWsChat handles websocket requests from peer
func serveWsChat(chatHub *ChatHub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &ChatClient{
		chatHub: chatHub,
		conn:    conn,
		send:    make(chan []byte, 256),
	}

	client.chatHub.register <- client

	// Allow collection of memory referenced by the caller by doing all work in new goroutines
	go client.writePump()
	go client.readPump()
}
