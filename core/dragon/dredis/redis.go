package dredis

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"time"
)

var (
	Redis *redis.Client
)

// init redis
func InitRedis() {
	timeout := viper.GetDuration("database.redis.timeout")
	Redis = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("database.redis.host") + ":" + viper.GetString("database.redis.port"),
		Password:     viper.GetString("database.redis.auth"), // password set
		DB:           viper.GetInt("database.redis.db"),      // use default DB
		ReadTimeout:  timeout * time.Millisecond,
		WriteTimeout: timeout * time.Millisecond,
		DialTimeout:  timeout * time.Millisecond,
		PoolSize:     300,
		MinIdleConns: 50,
	})
}
