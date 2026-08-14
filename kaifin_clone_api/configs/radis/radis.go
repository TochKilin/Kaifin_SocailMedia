package radis

import (
	// standard library
	"context"
	"encoding/json"
	"log"
	"sync"
	"time"

	// commmunty package

	"github.com/redis/go-redis/v9"

	"kaifin_clone_api/pkg/logs"
)

var (
	once   sync.Once
	client *redis.Client
)

// for catch seasoin & token
func NewRedisClient() *redis.Client {
	redis_config := InitRedis()
	// create redis client only one
	once.Do(func() {
		// connect redis serviver
		client = redis.NewClient(&redis.Options{
			Addr:     redis_config.RedisHost + ":" + redis_config.RedisPort,
			Password: redis_config.RedisPassword,
			DB:       redis_config.RedisDB,
		})
		// set redis context 3 second
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		// check redis live or not
		pong, err := client.Ping(ctx).Result()
		if err != nil {
			// log only — Fatalf
			logs.NewCustomLog("connect_redis_failed", err.Error(), "error")
			log.Printf("Warning: Redis unavailable: %v", err)
			return
		}
		log.Printf("Connected to Redis successfully: %s", pong)
	})

	return client
}

type RedisUtil struct {
	rdb *redis.Client
}

func NewRedisUtil(rdb *redis.Client) *RedisUtil {
	return &RedisUtil{rdb: rdb}
}

func (ru *RedisUtil) SetCacheKey(key string, value interface{}, ctx context.Context) error {
	// if redis not conn skips
	if ru.rdb == nil {
		log.Println("Warning: Redis client is nil, skipping cache store.")
		return nil
	}
	// convert struct to string
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	if err != nil {
		log.Printf("Warning: Redis Server is offline, cannot set key %s: %v", key, err)
		return nil
	}
	// set exspired 24h
	return ru.rdb.Set(ctx, key, data, 24*time.Hour).Err()
}
