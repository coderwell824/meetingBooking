package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"meetingBooking/pkg/format"
	"meetingBooking/services"
	"meetingBooking/utils"
	"net/http"
)

func CreateCategoryHandler(ctx *gin.Context) {
	createCategoryService := services.CategoryService{}
	if err := ctx.ShouldBind(&createCategoryService); err != nil {
		msg := utils.GetValidMsg(err, &createCategoryService)
		ctx.JSON(http.StatusBadRequest, format.RespErrorWithData(errors.New(msg)))
	} else {
		resp, respErr := createCategoryService.Create(ctx.Request.Context(), createCategoryService)
		if respErr != nil {
			ctx.JSON(http.StatusInternalServerError, format.RespErrorWithData(respErr))
		} else {
			ctx.JSON(http.StatusOK, format.RespSuccessWithData(resp))
		}
	}
}

func UpdateCategoryHandler(ctx *gin.Context) {
	updateCategoryService := services.UpdateCategoryService{}
	if err := ctx.ShouldBind(&updateCategoryService); err != nil {
		msg := utils.GetValidMsg(err, &updateCategoryService)
		ctx.JSON(http.StatusBadRequest, format.RespErrorWithData(errors.New(msg)))
	} else {
		resp, respErr := updateCategoryService.Update(ctx.Request.Context(), updateCategoryService)
		if respErr != nil {
			ctx.JSON(http.StatusInternalServerError, format.RespErrorWithData(respErr))
		} else {
			ctx.JSON(http.StatusOK, format.RespSuccessWithData(resp))
		}
	}
}
func CategoryListHandler(ctx *gin.Context) {
	categoryListService := services.CategoryService{}
	resp, respErr := categoryListService.List(ctx.Request.Context())
	if respErr != nil {
		ctx.JSON(http.StatusInternalServerError, format.RespErrorWithData(respErr))
	} else {
		ctx.JSON(http.StatusOK, format.RespSuccessWithData(resp))
	}
}
func DeleteCategoryHandler(ctx *gin.Context) {
	categoryId := ctx.Query("categoryId")
	if categoryId == "" {
		ctx.JSON(http.StatusInternalServerError, format.RespErrorWithData(errors.New("缺少分类id")))
		return
	}
	categoryDeleteService := services.CategoryService{}
	resp, respErr := categoryDeleteService.Delete(ctx.Request.Context(), categoryId)
	if respErr != nil {
		ctx.JSON(http.StatusInternalServerError, format.RespErrorWithData(respErr))
	} else {
		ctx.JSON(http.StatusOK, format.RespSuccessWithData(resp))
	}
}
