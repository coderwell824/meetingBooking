package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"meetingBooking/config"
	"meetingBooking/middleware"
	"meetingBooking/routes/module"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("something-very-secret"))
	r.Use(middleware.Cors())
	r.Use(sessions.Sessions("my-session", store))
	r.StaticFS("/static", http.Dir("./assets")) //获取服务器上的静态资源
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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
	module.LoadCategoriesRoute(authed)
	module.LoadProductRoute(authed)

	return r
}
