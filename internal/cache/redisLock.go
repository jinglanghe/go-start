package cache

import (
	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	"github.com/jinglanghe/go-start/internal/config"
)

var RedisLock *redsync.Redsync

func InitRedisLock() {
	redisConf := config.Config.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisConf.Addr,
		Password: redisConf.Password,
		DB:       redisConf.Db,
	})
	pool := goredis.NewPool(client)

	RedisLock = redsync.New(pool)
}
