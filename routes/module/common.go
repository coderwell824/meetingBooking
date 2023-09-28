package module

import (
	"github.com/gin-gonic/gin"
	"meetingBooking/api"
)

func LoadCommonRoute(v1 *gin.RouterGroup) {

	common := v1.Group("/")
	common.POST("file/upload", api.UploadFileHandler())
	common.GET("images/:filename", api.GetImageHandler())

}
