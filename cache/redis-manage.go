package cache

import "os"

var redisKeyWordTemp *redisBotTextCache
var redisImageMap *redisImageMapCache

func RedisInit() {
	redisKeyWordTemp = NewTextBotRedisCache(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PASSWORD"), 0)
	redisImageMap = NewImageMapRedisCache(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PASSWORD"), 0)
}

func GetRedisKeyWordTemp() *redisBotTextCache {
	return redisKeyWordTemp
}

func GetRedisImageMap() *redisImageMapCache {
	return redisImageMap
}
