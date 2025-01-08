package socket

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// WebSocket Upgrader
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func RegisterRoutes(r *gin.Engine) {
	r.GET("/ws/queue_manager", WebsocketHandler)
	r.GET("/ws/", WebsocketHandler)

}

// WebSocket handler
func WebsocketHandler(c *gin.Context) {
	w := c.Writer
	r := c.Request

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	for {
		// Read message from client
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}

		// Show message
		log.Printf("Received message: %s", message)
		userID, _ := c.Get("userID")
		log.Println(userID)

		userIDUint64, _ := strconv.ParseUint(fmt.Sprintf("%v", userID), 10, 32)
		userIDUint := uint(userIDUint64)
		fmt.Printf("🚀 ~ func SaveNotificationToken ~ userIDUint: %d, type: %T\n", userIDUint, userIDUint)

		// Send message to client
		err = conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println(err)
			break
		}
	}
}
