package main

import "encoding/json"

// ChatHub directs client traffic via different channels
type ChatHub struct {
	//Registered clients
	clients map[*ChatClient]bool

	rooms map[string]map[*ChatClient]bool

	//Inbound messages from the clients
	broadcast chan []byte

	//Register requests from the clients
	register chan *ChatClient

	//Unregister requests from clients
	unregister chan *ChatClient
}

func makeHub() *ChatHub {
	return &ChatHub{
		broadcast:  make(chan []byte),
		register:   make(chan *ChatClient),
		unregister: make(chan *ChatClient),
		clients:    make(map[*ChatClient]bool),
		rooms:      make(map[string]map[*ChatClient]bool),
	}
}

func sendToClientMsg(client *ChatClient, category string, text string) {
	currMsg := ToClientMsg{Cat: category, Text: text}
	cbytes, err2 := json.Marshal(currMsg)
	if err2 != nil {
		panic(err2)
	}
	client.send <- cbytes
}

func (chatHub *ChatHub) run() {
	for {
		select {
		case client := <-chatHub.register:
			// Receive instances of ChatClient to register
			chatHub.clients[client] = true
			chatHub.addToRoom(client, client.room)
		case client := <-chatHub.unregister:
			// Only unregister client if it's in our clients map of ChatHub
			if _, ok := chatHub.clients[client]; ok {
				close(client.send)
				delete(chatHub.clients, client)
			}
		case message := <-chatHub.broadcast:
			for client := range chatHub.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(chatHub.clients, client)
				}
			}
		}
	}
}
