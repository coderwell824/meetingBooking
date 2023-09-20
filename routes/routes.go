package routes

import (
	"github.com/gin-gonic/gin"
	"meetingBooking/config"
	"meetingBooking/routes/module"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	//store := sessions.NewCookieStore([]byte("something-very-secret"))
	//开始swag
	//r.Use(sessions.Sessions("mysession", store))
	//r.Use(middleware.Cors())

	v1 := r.Group(config.BaseUrl)

	v1.GET("ping", func(context *gin.Context) {
		context.JSON(200, "success")
	})

	module.LoadUserRoute(v1)
	module.LoadRoomsRoute(v1)
	//_ := v1.Group("/") //登录保护
	//authed.Use(middleware.JWT())

	return r
}