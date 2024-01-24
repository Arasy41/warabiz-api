package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

/**
** Description
*? key:		redis key
*? data:	something you want to keep in a redis
*? exp:		redis key life time
 */

// Custom function to store data using redis Set
func RedisSet(ctx context.Context, client redis.Cmdable, key string, data interface{}, exp time.Duration) error {
	op1 := client.Set(ctx, key, data, exp)
	if err := op1.Err(); err != nil {
		return fmt.Errorf("Unable to SET data. error: %v", err)
	}
	return nil
}

// Custom function to get data using redis Get
func RedisGet(ctx context.Context, client redis.Cmdable, key string) (string, error) {
	op2 := client.Get(ctx, key)
	if err := op2.Err(); err != nil {
		return "", fmt.Errorf("Unable to GET data. error: %v", err)
	}
	res, err := op2.Result()
	if err != nil {
		return "", fmt.Errorf("Unable to GET data. error: %v", err)
	}
	return res, nil
}

// Custom function to delete key using redis Del
func RedisDel(ctx context.Context, client redis.Cmdable, key string) error {
	op2 := client.Del(ctx, key)
	if err := op2.Err(); err != nil {
		return fmt.Errorf("Unable to DELETE data. error: %v", err)
	}
	return nil
}

func RedisGetTTl(ctx context.Context, client redis.Cmdable, key string) (*time.Duration, error) {
	expiration, err := client.TTL(ctx, key).Result()
	if err != nil {
		return nil, fmt.Errorf("Unable to GET data. error: %v", err)
	}
	return &expiration, nil
}
