package certmon

import (
	"github.com/go-redis/redis"
	"github.com/json-iterator/go"
	"os"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func CreateRedisClient() *redis.Client {
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")

	if len(redisHost) == 0 {
		redisHost = "localhost"
	}

	if len(redisPort) == 0 {
		redisPort = "6379"
	}

	return redis.NewClient(&redis.Options{
		Addr:     redisHost + ":" + redisPort,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
