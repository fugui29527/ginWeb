package util

import (
	"fmt"
	"ginWeb/config"
	"github.com/go-redis/redis"
	"time"
)

var redisClient *redis.Client

func CreateRedis() {
	redisInfo := config.AppConfig.RedisInfo
	redisClient = redis.NewClient(
		&redis.Options{
			Addr:     redisInfo.Addr + ":" + redisInfo.Port,
			Password: redisInfo.Password, // no password set
			DB:       redisInfo.DB,       // use default DB
		})
}

//设置过期时间
func SetNx(key string, val interface{}, expiration time.Duration) bool {
	bool, err := redisClient.SetNX(key, val, expiration).Result()
	if err != nil {
		fmt.Println("redis 保存错误:", err)
		return bool
	}
	return bool
}

//不设置过期时间
func Set(key string, val interface{}) string {
	string, err := redisClient.Set(key, val, 0).Result()
	if err != nil {
		fmt.Println("redis 保存错误:", err)
		return string
	}
	return string
}
