package module

import (
	"github.com/gin-gonic/gin"
	"meetingBooking/api"
)

func LoadPermissionsRoute(v1 *gin.RouterGroup) {

	roles := v1.Group("/permissions")

	roles.POST("create", api.CreatePermissionHandler())
	//
	//roles.POST("create", func(context *gin.Context) {
	//	context.JSON(200, "success")
	//})

	roles.GET("history", func(context *gin.Context) {
		context.JSON(200, "success")
	})

	roles.POST("approve", func(context *gin.Context) {
		context.JSON(200, "success")
	})

	roles.GET("cancel/:id", func(context *gin.Context) {
		context.JSON(200, "success")
	})

	roles.GET("unbind:id", func(context *gin.Context) {
		context.JSON(200, "success")
	})

	roles.GET("search", func(context *gin.Context) {
		context.JSON(200, "success")
	})

	roles.POST("urge", func(context *gin.Context) {
		context.JSON(200, "success")
	})

}
