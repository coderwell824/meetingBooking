package module

import (
	"github.com/gin-gonic/gin"
	"meetingBooking/api"
)

func LoadProductRoute(v1 *gin.RouterGroup) {

	product := v1.Group("/product")

	product.POST("create", api.CreateProductHandler)
	product.GET("list", api.ProductListHandler)
	product.PUT("/update/:productId", api.ProductUpdateHandler)

}
