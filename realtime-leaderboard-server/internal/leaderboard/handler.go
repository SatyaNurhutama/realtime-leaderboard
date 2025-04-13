package leaderboard

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ScoreRequest struct {
	User  string `json:"user"`
	Score int64  `json:"score"`
}

func SubmitScoreHandler(c *gin.Context) {
	var req ScoreRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := SubmitScore(req.User, req.Score); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Score submitted successfully"})
}

func GetTopScoresHandler(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil || limit <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit parameter"})
		return
	}
	scores, err := GetTopScores(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	result := make([]gin.H, 0, len(scores))
	for _, z := range scores {
		result = append(result, gin.H{
			"user":  z.Member,
			"score": z.Score,
		})
	}
	c.JSON(http.StatusOK, result)
}
