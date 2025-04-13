package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/satyanurhutama/realtime-leaderboard/internal/leaderboard"
	"github.com/satyanurhutama/realtime-leaderboard/pkg/redis_client"
)

func main() {
	redis_client.InitRedis()

	go leaderboard.RunHub()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Origin", "Content-Type", "Upgrader"},
	}))

	r.POST("/submit-score", leaderboard.SubmitScoreHandler)
	r.GET("/top-scores", leaderboard.GetTopScoresHandler)
	r.GET("/ws", leaderboard.WebSocketHandler)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
