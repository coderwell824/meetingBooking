package module

import (
	"github.com/gin-gonic/gin"
	"meetingBooking/api"
)

func LoadCommonRoute(v1 *gin.RouterGroup) {

	common := v1.Group("/")
	common.POST("file/upload", api.UploadFileHandler())
	common.GET("images/:filename", api.GetImageHandler())
	common.POST("addPos", api.AddPosHandler())
	common.GET("pos/:name", api.GetPosHandler())
	common.GET("pos/all", api.GetPosAllHandler())
	common.GET("nearbySearch", api.GetNearbySearchHandler())
}
