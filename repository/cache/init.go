package cache

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"meetingBooking/config"
	"strconv"
	"time"
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

func RedisGetKey(key string) (string, error) {
	if val, err := RedisClient.Get(key).Result(); err != nil {
		return "", err
	} else {
		return val, nil
	}
}

func RedisSetKey(key string, val interface{}, ttl time.Duration) error {
	if _, err := RedisClient.Set(key, val, ttl).Result(); err != nil {
		return err
	}
	return nil
}

//geoadd：增加某个位置的坐标
//geopos：获取某个位置的坐标。
//geohash：获取某个位置的geohash值
//geodist：获取两个位置的距离。
//georadius：根据给定位置坐标获取指定范围内的位置集合
//georadiusbymember：根据给定位置获取指定范围内的位置集合

func RedisGeoAdd(positionInfo ...*redis.GeoLocation) (err error) {
	err = RedisClient.GeoAdd("location", positionInfo...).Err()
	if err != nil {
		log.Println("redis位置信息设置失败")
		return
	}
	return nil
}

func RedisGetGeo(positionName ...string) (resPos []*redis.GeoPos, err error) {
	resPos, err = RedisClient.GeoPos("location", positionName...).Result()
	if err != nil {
		log.Println("redis位置信息获取失败")
		return
	}
	return
}
func RedisGeoRadius(longitude, latitude float64, radius float64) (resRadius []redis.GeoLocation, err error) {
	resRadius, err = RedisClient.GeoRadius("location", longitude, latitude, &redis.GeoRadiusQuery{
		Radius:      radius, //radius表示范围距离，
		Unit:        "km",   //距离单位是 m|km|ft|mi
		WithCoord:   true,   //传入WITHCOORD参数，则返回结果会带上匹配位置的经纬度
		WithDist:    true,   //传入WITHDIST参数，则返回结果会带上匹配位置与给定地理位置的距离。
		WithGeoHash: true,   //传入WITHHASH参数，则返回结果会带上匹配位置的hash值。
		Sort:        "ASC",  //默认结果是未排序的，传入ASC为从近到远排序，传入DESC为从远到近排序。
	}).Result()
	if err != nil {
		log.Println("redis位置返回信息获取失败")
		return
	}
	return
}
