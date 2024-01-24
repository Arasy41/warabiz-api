package redis

import (
	"warabiz/api/config"

	"github.com/go-redis/redis/v8"
)

// * Function to init new redis client
func NewRedisClient(cfg *config.RedisAccount) *redis.Client {
	redisHost := cfg.Host
	if redisHost == "" {
		redisHost = ":6379"
	}
	client := redis.NewClient(&redis.Options{
		Addr:         redisHost,
		Password:     cfg.Password,
		MinIdleConns: cfg.MinIdleConns,
		PoolSize:     cfg.PoolSize,
		PoolTimeout:  cfg.PoolTimeout,
		DB:           cfg.DB,
	})
	return client
}
