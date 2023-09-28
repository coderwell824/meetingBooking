package routes

import (
	"github.com/gin-gonic/gin"
	"meetingBooking/config"
	"meetingBooking/middleware"
	"meetingBooking/routes/module"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	//store := sessions.NewCookieStore([]byte("something-very-secret"))
	//开始swag
	//r.Use(sessions.Sessions("mysession", store))
	r.Use(middleware.Cors())

	v1 := r.Group(config.BaseUrl)

	v1.GET("ping", func(context *gin.Context) {
		context.JSON(200, "success")
	})

	module.LoadLoginRoute(v1)
	module.LoadCommonRoute(v1)
	authed := v1.Group("/") //登录保护
	authed.Use(middleware.JWT())

	module.LoadPermissionsRoute(authed)
	module.LoadUserRoute(authed)
	module.LoadBookingsRoute(authed)
	module.LoadRolesRoute(authed)
	module.LoadRoomsRoute(authed)
	module.LoadStatisticsRoute(authed)

	return r
}
