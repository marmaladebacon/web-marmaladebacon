package main

import "fmt"

func (chatHub *ChatHub) addToRoom(c *ChatClient, roomName string) {
	_, roomPresent := chatHub.rooms[roomName]
	if !roomPresent {
		chatHub.rooms[roomName] = make(map[*ChatClient]bool)
	}
	chatHub.rooms[roomName][c] = true
	sendToClientMsg(c, "admin", fmt.Sprintf("Joining: %v", roomName))
}

func (chatHub *ChatHub) removeFromRoom(c *ChatClient, roomName string) {
	_, roomPresent := chatHub.rooms[roomName]
	if roomPresent {
		delete(chatHub.rooms[roomName], c)
	}
}
