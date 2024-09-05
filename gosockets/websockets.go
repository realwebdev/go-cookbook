package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var wsUpgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := wsUpgrade.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to set websocket upgrade: %+v", err)
		return
	}

	for {
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Failed to read message: %+v", err)
			break
		}

		msg = []byte("pong")

		conn.WriteMessage(messageType, msg)
	}
}

func main() {
	r := gin.Default()

	r.LoadHTMLFiles("index.html")
	r.GET("/ping", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/ws", func(c *gin.Context) {
		wsHandler(c.Writer, c.Request)
	})

	r.Run("localhost:8080")
}
