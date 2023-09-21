package module

import (
	"github.com/gin-gonic/gin"
	"meetingBooking/api"
)

func LoadUserRoute(v1 *gin.RouterGroup) {
	
	user := v1.Group("/user")
	
	user.POST("register", api.UserRegisterHandler())
	user.POST("login", func(context *gin.Context) {
		context.JSON(200, "success")
	})
	
	user.POST("info/update", func(context *gin.Context) {
		context.JSON(200, "success")
	})
	
	user.POST("updatePassword", func(context *gin.Context) {
		context.JSON(200, "success")
	})
	
	user.GET("list", func(context *gin.Context) {
		context.JSON(200, "success")
	})
	
	user.GET("freeze", func(context *gin.Context) {
		context.JSON(200, "success")
	})
	
}
