package main

import "encoding/json"

// ChatHub directs client traffic via different channels
type ChatHub struct {
	//Registered clients
	clients map[*ChatClient]bool

	rooms map[string]map[*ChatClient]bool

	//Inbound messages from the clients
	broadcast chan *BroadcastMsg

	//Register requests from the clients
	register chan *ChatClient

	//Unregister requests from clients
	unregister chan *ChatClient
}

// BroadcastMsg is for sending messages to the broadcast pipe
type BroadcastMsg struct {
	client   *ChatClient
	category string
	text     []byte
}

func makeHub() *ChatHub {
	return &ChatHub{
		broadcast:  make(chan *BroadcastMsg),
		register:   make(chan *ChatClient),
		unregister: make(chan *ChatClient),
		clients:    make(map[*ChatClient]bool),
		rooms:      make(map[string]map[*ChatClient]bool),
	}
}

func makeMsgFromClient(client *ChatClient, text string) *BroadcastMsg {
	return &BroadcastMsg{
		client: client,
		text:   []byte(text),
	}
}

func makeToClientMsg(b *BroadcastMsg) ToClientMsg {
	if b.client != nil {
		return ToClientMsg{Cat: "::", Text: string(b.text)}
	}

	return ToClientMsg{Cat: b.category, Text: string(b.text)}
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
				temp := makeToClientMsg(message)
				sendToClientMsg(client, temp.Cat, temp.Text)
				/*
					select {
					case client.send <- makeToClientMsg(message):
					default:
						close(client.send)
						delete(chatHub.clients, client)
					}*/
			}
		}
	}
}
