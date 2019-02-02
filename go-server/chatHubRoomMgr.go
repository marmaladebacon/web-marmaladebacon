package main

import (
	"bytes"
	"fmt"
	"strings"
)

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

func handleMsg(message []byte, chatClient *ChatClient) {
	s := string(message[:])

	sArray := strings.Split(s, " ")
	firstChar := string(sArray[0][0])

	handler := broadcastUserMsgToRoom
	if firstChar == "/" {
		handler = handleUserCommand
	}
	handler(message, chatClient)
}

func handleUserCommand(msg []byte, chatClient *ChatClient) {
	s := string(msg[0:])
	sArray := strings.Split(s, " ")
	cmd := strings.ToLower(string(sArray[0][1:]))
	fmt.Printf("Command %v", cmd)

	switch cmd {
	case "help":

	}
}

func broadcastUserMsgToRoom(msg []byte, chatClient *ChatClient) {
	// replace all newlines with space,
	// then remove all leading and trailing whitespace
	textMessage := bytes.TrimSpace(bytes.Replace(msg, newline, space, -1))

	broadcastMsg := BroadcastMsg{client: chatClient, text: textMessage}

	// broadcast user message to the hub
	chatClient.chatHub.broadcast <- &broadcastMsg
}

/* Help responders */

func defaultError(chatClient *ChatClient) {
	sendToClientMsg(chatClient, "admin", "Unknown command, try using /help")
}
