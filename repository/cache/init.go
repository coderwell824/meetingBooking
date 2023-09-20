package cache

import (
	"fmt"
	"github.com/go-redis/redis"
	"meetingBooking/config"
	"strconv"
)

var RedisClient *redis.Client

// RedisInit 初始化Redis
func RedisInit() {
	db, _ := strconv.ParseUint(config.RedisDbName, 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPw,
		DB:       int(db),
	})
	_, err := client.Ping().Result()
	
	if err != nil {
		panic(err.Error())
	}
	
	RedisClient = client
	fmt.Println("redis 连接")
}
