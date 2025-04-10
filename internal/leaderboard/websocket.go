package leaderboard

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/satyanurhutama/realtime-leaderboard/pkg/redis_client"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan map[string]interface{})
var upgrader = websocket.Upgrader{}

func WebSocketHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.String(500, "Failed to upgrade connection: %v", err)
		return
	}
	defer conn.Close()
	clients[conn] = true
}

func RunHub() {
	pubsub := redis_client.GetRedisClient().Subscribe(redis_client.GetCtx(), "score_updates")
	ch := pubsub.Channel()

	go func() {
		for msg := range ch {
			var data map[string]interface{}
			_ = json.Unmarshal([]byte(msg.Payload), &data)
			broadcast <- data
		}
	}()

	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				client.Close()
				delete(clients, client)
			}
		}
	}
}
