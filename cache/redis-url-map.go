package cache

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis"
)

type redisImageMapCache struct {
	host     string
	db       int
	expires  time.Duration
	password string
}

func NewImageMapRedisCache(host string, password string, db int) *redisImageMapCache {
	return &redisImageMapCache{
		host:     host,
		db:       db,
		expires:  30 * time.Minute,
		password: password,
	}
}

func (cache *redisImageMapCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		DB:       cache.db,
		Password: cache.password,
	})
}

func (cache *redisImageMapCache) Set(key string, value map[string]string) {
	json, _ := json.Marshal(value)
	cache.getClient().Set(imageTempKey(key), json, cache.expires)
}

func (cache *redisImageMapCache) Get(key string) map[string]string {
	m := make(map[string]string)
	val, _ := cache.getClient().Get(imageTempKey(key)).Bytes()
	json.Unmarshal(val, &m)
	return m
}

func (cache *redisImageMapCache) Del(key string) {
	cache.getClient().Del(imageTempKey(key))
}

func imageTempKey(key string) string {
	return key + "-imageTemp"
}
