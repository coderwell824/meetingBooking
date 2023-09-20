package main

import (
	"meetingBooking/config"
	"meetingBooking/repository/cache"
	"meetingBooking/repository/db/dao"
)

func main() {
	loadingConfig()
}

func loadingConfig() {
	config.Init()
	dao.MySqlInit()
	cache.RedisInit()
}
