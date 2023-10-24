package main

import (
	"meetingBooking/config"
	_ "meetingBooking/docs"
	"meetingBooking/repository/cache"
	"meetingBooking/repository/db/dao"
	"meetingBooking/repository/mongo"
	"meetingBooking/routes"
	"meetingBooking/services/ws"
)

// @title go_server API文档
// @version 1.0
// @description API文档
// @host 127.0.0.01:8888
// @BasePath /api/v1
func main() {
	loadingConfig()
	r := routes.NewRouter()
	_ = r.Run(config.HttpPort)

}

func loadingConfig() {
	config.Init()
	dao.InitMysqlConnection()
	go ws.Manager.Watch() //开启websocket通信
	cache.InitRedisConnection()
	mongo.InitMongoConnection()
}
