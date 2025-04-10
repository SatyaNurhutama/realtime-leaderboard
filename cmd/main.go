package main

import (
	"github.com/gin-gonic/gin"
	"github.com/satyanurhutama/realtime-leaderboard/internal/leaderboard"
	"github.com/satyanurhutama/realtime-leaderboard/pkg/redis_client"
)

func main() {
	redis_client.InitRedis()

	r := gin.Default()
	r.POST("/submit-score", leaderboard.SubmitScoreHandler)
	r.GET("/top-scores", leaderboard.GetTopScoresHandler)
	r.GET("/ws", leaderboard.WebSocketHandler)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
