package cache

import (
	"context"
	"github.com/jinglanghe/go-start/internal/config"
	"github.com/jinglanghe/go-start/utils/log"
	"time"

	"github.com/go-redis/redis/v8"
)

// Redis cache implement
type Redis struct {
	ctx         context.Context
	client      *redis.Client
	redisExpire int32
}

// Connect connect to redis
func (r *Redis) Connect() {
	r.ctx = context.Background()
	redisConf := config.Config.Redis
	r.client = redis.NewClient(&redis.Options{
		Addr:     redisConf.Addr,
		Password: redisConf.Password,
		DB:       redisConf.Db,
	})
	_, err := r.client.Ping(r.ctx).Result()
	if err != nil {
		log.Fatal().Err(err).Msg("Could not connected to redis")
	}
	log.Info().Msg("Successfully connected to redis")
	r.redisExpire = int32(redisConf.Expire / time.Second)
}

func (r *Redis) Close() {
	r.client.Close()
}

// Get from key
func (r *Redis) Get(key string) (string, error) {
	return r.client.Get(r.ctx, key).Result()
}

// Set value with key and expire time
func (r *Redis) Set(key string, val string, expire int) error {
	return r.client.Set(r.ctx, key, val, time.Duration(expire)).Err()
}

// Del delete key in redis
func (r *Redis) Del(key string) error {
	return r.client.Del(r.ctx, key).Err()
}

// HashGet from key
func (r *Redis) HashGet(hk, key string) (string, error) {
	return r.client.HGet(r.ctx, hk, key).Result()
}

// HashDel delete key in specify redis's hashtable
func (r *Redis) HashDel(hk, key string) error {
	return r.client.HDel(r.ctx, hk, key).Err()
}

// Increase increase value
func (r *Redis) Increase(key string) error {
	return r.client.Incr(r.ctx, key).Err()
}

// Expire set ttl
func (r *Redis) Expire(key string, dur time.Duration) error {
	return r.client.Expire(r.ctx, key, dur).Err()
}

func (r *Redis) SMembers(key string) ([]string, error) {
	return r.client.SMembers(r.ctx, key).Result()
}

func (r *Redis) HashGetAll(hk string) (map[string]string, error) {
	return r.client.HGetAll(r.ctx, hk).Result()
}

func (r *Redis) HashSet(hk, key string, value interface{}) error {
	return r.client.HSet(r.ctx, hk, key, value).Err()
}

func (r *Redis) SAdd(key string, member string) error {
	return r.client.SAdd(r.ctx, key, member).Err()
}

func (r *Redis) SRem(key string, member string) error {
	return r.client.SRem(r.ctx, key, member).Err()
}

func (r *Redis) HashMSet(hk string, params map[string]interface{}) error {
	return r.client.HMSet(r.ctx, hk, params).Err()
}
