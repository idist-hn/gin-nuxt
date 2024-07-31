package redisProvider

import (
	"github.com/go-redis/redis"
	"idist-core/app/providers/configProvider"
	"idist-core/app/providers/loggerProvider"
)

var client *redis.Client

func Init() {
	loggerProvider.GetLogger().Info("------------------------------------------------------------")
	cf := configProvider.GetConfig()
	client = redis.NewClient(&redis.Options{
		Addr:     cf.GetString("redis.addr"),
		Password: cf.GetString("redis.password"),
		DB:       0,
	})

	if cf.GetString("env") == "development" {
		if message, err := client.FlushAll().Result(); err != nil {
			loggerProvider.GetLogger().Info(message)
		} else {
			loggerProvider.GetLogger().Info("Redis flushed all cache.")
		}
	}

	if _, err := client.Ping().Result(); err != nil {
		loggerProvider.GetLogger().Info("Redis server can't connect.")
	} else {
		loggerProvider.GetLogger().Info("Redis server connected.")
	}
}

func GetClient() *redis.Client {
	return client
}
