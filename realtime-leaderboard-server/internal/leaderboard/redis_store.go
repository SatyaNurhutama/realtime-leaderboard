package leaderboard

import (
	"encoding/json"

	"github.com/redis/go-redis/v9"
	"github.com/satyanurhutama/realtime-leaderboard/pkg/redis_client"
)

const leaderboardKey = "leaderboard"

func SubmitScoreToRedis(user string, score int64) (float64, error) {
	newScore, err := redis_client.GetRedisClient().ZIncrBy(redis_client.GetCtx(), leaderboardKey, float64(score), user).Result()
	return newScore, err
}

func GetTopNFromRedis(n int64) ([]redis.Z, error) {
	return redis_client.GetRedisClient().ZRevRangeWithScores(redis_client.GetCtx(), leaderboardKey, 0, n-1).Result()
}

func PublishScoreUpdate(user string, score int64) error {
	msg, _ := json.Marshal(map[string]any{
		"user":  user,
		"score": score,
	})
	return redis_client.GetRedisClient().Publish(redis_client.GetCtx(), "score_updates", msg).Err()
}
