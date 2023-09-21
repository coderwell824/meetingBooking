package module

import "github.com/gin-gonic/gin"

func LoadBookingsRoute(v1 *gin.RouterGroup) {
	
	bookings := v1.Group("/bookings")
	
	bookings.GET("list", func(context *gin.Context) {
		context.JSON(200, "success")
	})
	
	bookings.POST("create", func(context *gin.Context) {
		context.JSON(200, "success")
	})
	
	bookings.GET("history", func(context *gin.Context) {
		context.JSON(200, "success")
	})
	
	bookings.POST("approve", func(context *gin.Context) {
		context.JSON(200, "success")
	})
	
	bookings.GET("cancel/:id", func(context *gin.Context) {
		context.JSON(200, "success")
	})
	
	bookings.GET("unbind:id", func(context *gin.Context) {
		context.JSON(200, "success")
	})
	
	bookings.GET("search", func(context *gin.Context) {
		context.JSON(200, "success")
	})
	
	bookings.POST("urge", func(context *gin.Context) {
		context.JSON(200, "success")
	})
	
}
