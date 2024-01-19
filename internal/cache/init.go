package cache

import (
	"github.com/jinglanghe/go-start/utils/log"
	"time"
)

var adapter Adapter

type Adapter interface {
	Connect()
	Get(key string) (string, error)
	Set(key string, val string, expire int) error
	Del(key string) error
	HashGet(hk, key string) (string, error)
	HashDel(hk, key string) error
	Increase(key string) error
	Expire(key string, dur time.Duration) error
	SMembers(key string) ([]string, error)
	HashGetAll(hk string) (map[string]string, error)
	HashSet(hk, key string, value interface{}) error
	SAdd(key string, member string) error
	SRem(key string, member string) error
	HashMSet(hk string, pairs map[string]interface{}) error
	Close()
}

func Init() (Adapter, func()) {
	adapter = &Redis{}
	adapter.Connect()

	clearFunc := func() {
		if adapter != nil {
			adapter.Close()
			log.Info().Msg("closed cache connection")
		}
	}
	InitRedisLock()
	return adapter, clearFunc
}

// Set val in cache
func Set(key, val string, expire int) error {
	return adapter.Set(key, val, expire)
}

// Get val in cache
func Get(key string) (string, error) {
	return adapter.Get(key)
}

// Del delete key in cache
func Del(key string) error {
	return adapter.Del(key)
}

// HashGet get val in hashtable cache
func HashGet(hk, key string) (string, error) {
	return adapter.HashGet(hk, key)
}

// HashDel delete one key:value pair in hashtable cache
func HashDel(hk, key string) error {
	return adapter.HashDel(hk, key)
}

// Increase value
func Increase(key string) error {
	return adapter.Increase(key)
}

func Expire(key string, dur time.Duration) error {
	return adapter.Expire(key, dur)
}

func SMembers(key string) ([]string, error) {
	return adapter.SMembers(key)
}

func HashGetAll(hk string) (map[string]string, error) {
	return adapter.HashGetAll(hk)
}

func HashSet(hk, key string, value interface{}) error {
	return adapter.HashSet(hk, key, value)
}

func SAdd(key string, member string) error {
	return adapter.SAdd(key, member)
}

func SRem(key string, member string) error {
	return adapter.SRem(key, member)
}

func HashMSet(hk string, value map[string]interface{}) error {
	return adapter.HashMSet(hk, value)
}
