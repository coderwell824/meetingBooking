package module

import "github.com/gin-gonic/gin"

func LoadRoomsRoute(v1 *gin.RouterGroup) {

	rooms := v1.Group("/rooms")

	rooms.GET("list", func(context *gin.Context) {
		context.JSON(200, "success")
	})
	rooms.DELETE("delete/:id", func(context *gin.Context) {
		context.JSON(200, "success")
	})

	rooms.PUT("update/:id", func(context *gin.Context) {
		context.JSON(200, "success")
	})

	rooms.POST("create", func(context *gin.Context) {
		context.JSON(200, "success")
	})

	rooms.GET("search", func(context *gin.Context) {
		context.JSON(200, "success")
	})

}
