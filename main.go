package main

import (
	"meetingBooking/config"
	"meetingBooking/repository/cache"
	"meetingBooking/repository/db/dao"
	"meetingBooking/routes"
)

func main() {
	loadingConfig()
	r := routes.NewRouter()
	_ = r.Run(config.HttpPort)

}

func loadingConfig() {
	config.Init()
	dao.MySqlInit()
	cache.RedisInit()
}
