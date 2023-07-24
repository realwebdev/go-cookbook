package main

import (
	"fmt"
	"net/http"

	"websocket"
)

func websocketHandler(ws *websocket.Conn) {
	for {
		// read a message from web socket connection
		messageType, message, err := ws.ReadMessage()
		if err != nil {
			break
		}

		// print message
		fmt.Println(message)

		// send message back to websocket connection
		err = ws.WriteMessage(messageType, message)
		if err != nil {
			break
		}
	}
}

func main() {
	// listen for websocket connections on port 8080
	http.HandleFunc("/websocket", websocketHandler)
}
