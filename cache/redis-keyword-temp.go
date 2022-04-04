package cache

import (
	"encoding/json"
	"go-meme-bot/model"
	"time"

	"github.com/go-redis/redis"
)

type redisBotTextCache struct {
	host     string
	db       int
	expires  time.Duration
	password string
}

func NewTextBotRedisCache(host string, password string, db int) *redisBotTextCache {
	return &redisBotTextCache{
		host:     host,
		db:       db,
		expires:  60 * time.Second,
		password: password,
	}
}

func (cache *redisBotTextCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		DB:       cache.db,
		Password: cache.password,
	})
}

func (cache *redisBotTextCache) Set(key string, value model.UserKeyWord) {
	json, _ := json.Marshal(value)
	cache.getClient().Set(key, json, cache.expires)
}

func (cache *redisBotTextCache) Get(key string) *model.UserKeyWord {
	user := &model.UserKeyWord{}
	val, _ := cache.getClient().Get(key).Bytes()
	json.Unmarshal(val, &user)
	return user
}

func (cache *redisBotTextCache) Del(key string) {
	cache.getClient().Del(key)
}

