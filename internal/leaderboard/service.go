package leaderboard

import "github.com/redis/go-redis/v9"

func SubmitScore(user string, score int64) error {
	if err := SubmitScoreToRedis(user, score); err != nil {
		return err
	}
	return PublishScoreUpdate(user, score)
}

func GetTopScores(limit int64) ([]redis.Z, error) {
	return GetTopNFromRedis(limit)
}
