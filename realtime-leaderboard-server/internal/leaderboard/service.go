package leaderboard

import "github.com/redis/go-redis/v9"

func SubmitScore(user string, score int64) error {
	newScore, err := SubmitScoreToRedis(user, score)
	if err != nil {
		return err
	}
	return PublishScoreUpdate(user, int64(newScore))
}

func GetTopScores(limit int64) ([]redis.Z, error) {
	return GetTopNFromRedis(limit)
}
