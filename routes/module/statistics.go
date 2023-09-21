package module

import "github.com/gin-gonic/gin"

func LoadStatisticsRoute(v1 *gin.RouterGroup) {
	
	rooms := v1.Group("/statistics")
	
	rooms.GET("meetingRoom", func(context *gin.Context) {
		context.JSON(200, "success")
	})
	rooms.GET("userBooking", func(context *gin.Context) {
		context.JSON(200, "success")
	})
	
}
