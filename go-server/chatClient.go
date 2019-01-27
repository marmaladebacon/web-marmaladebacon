package main

import (
	"bytes"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer, 10 seconds
	writeWait = 10 * time.Second
	// Time allowed to read the next pong message form the peer
	pongWait = 60 * time.Second
	// Send pings to peer with this period. Must be less than pongWait
	pingPeriod = (pongWait * 9) / 10
	//Maximum message size allowed from peer
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  2048,
	WriteBufferSize: 2048,
}

// ToClientMsg is a struct that contains text and metaData for the chat
type ToClientMsg struct {
	Cat  string `json:"category"`
	Text string `json:"text"`
}

// ChatClient handles websocket client connections
type ChatClient struct {
	room    string
	chatHub *ChatHub
	conn    *websocket.Conn

	// Buffered channel of outBoundMessages
	send chan []byte
	// Buffered channel of new ToClientMsg
	sendClientMsg chan ToClientMsg
}

// readPump pumps messages from the websocket connection to the hub
// eg user types something, received by the client model on the server
// and this message gets sent to the hub in here
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one read on a connection by executing all
// reads from this goRoutine
func (client *ChatClient) readPump() {
	defer func() {
		// When forever look is broken out of (see below),
		// unregister this client from chatHub and close the connection)
		client.chatHub.unregister <- client
		client.conn.Close()
	}()

	client.conn.SetReadLimit(maxMessageSize)
	client.conn.SetReadDeadline(time.Now().Add(pongWait))
	client.conn.SetPongHandler(func(string) error {
		client.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})
	for {
		//Read the message being sent from the client
		_, message, err := client.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		// replace all newlines with space,
		// then remove all leading and trailing whitespace
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))

		// broadcast user message to the hub
		client.chatHub.broadcast <- message
	}
}

//writePump pumps messages from the hub to the websocket connection
//
// A goroutine running writePump is started for each connection.
// The application ensures there is at most one write to a connection by executing
// all writes from this goroutine
func (client *ChatClient) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		client.conn.Close()
	}()

	for {
		select {
		case message, ok := <-client.send:
			client.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				//The chathub has closed the channel
				client.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := client.conn.NextWriter(websocket.BinaryMessage)
			if err != nil {
				return
			}

			w.Write(message)

			// Add queued chat messages to the current websocket message
			n := len(client.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-client.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			client.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := client.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				//if err is not nil we return?
				return
			}
		}
	}
}
