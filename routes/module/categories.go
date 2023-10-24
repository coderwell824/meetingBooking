package module

import (
	"github.com/gin-gonic/gin"
	"meetingBooking/api"
)

func LoadCategoriesRoute(v1 *gin.RouterGroup) {

	category := v1.Group("/category")
	category.POST("/create", api.CreateCategoryHandler)
	category.PUT("/update", api.UpdateCategoryHandler)
	category.GET("/list", api.CategoryListHandler)
	category.DELETE("/delete", api.DeleteCategoryHandler)

}
