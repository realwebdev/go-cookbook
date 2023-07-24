package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	// upgrading connection to websocket
	conn, err := websocket.Upgrade(w, r, nil, 1024, 1024)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Listen for messages from the client
	for {
		// Read a message from the client
		msgType, _, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}

		// echo the message back to the client
		conn.WriteMessage(msgType, []byte("Hello from the server!"))
	}
}

func main() {
	http.HandleFunc("/ws", wsEndpoint)
	http.ListenAndServe(":8080", nil)
}
