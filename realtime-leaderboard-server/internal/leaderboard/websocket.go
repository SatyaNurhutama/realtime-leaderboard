package leaderboard

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/satyanurhutama/realtime-leaderboard/pkg/redis_client"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan map[string]interface{})
var upgrader = websocket.Upgrader{}
var lock = sync.Mutex{}

func WebSocketHandler(c *gin.Context) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.String(500, "Failed to upgrade connection: %v", err)
		return
	}

	lock.Lock()
	clients[conn] = true
	lock.Unlock()

	defer func() {
		lock.Lock()
		delete(clients, conn)
		lock.Unlock()
		conn.Close()
	}()

	// Read loop just to keep connection alive
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			break
		}
	}
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

		lock.Lock()
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				client.Close()
				delete(clients, client)
			}
		}
		lock.Unlock()
	}
}
