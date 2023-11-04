package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"meetingBooking/pkg/format"
	"meetingBooking/services"
	"meetingBooking/utils"
	"net/http"
)

// CreateProductHandler 新增商品
func CreateProductHandler(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, format.RespErrorWithData(errors.New("缺少商品图片")))
		return
	}
	files := form.File["file"]
	createProductService := services.ProductService{}
	if err := ctx.ShouldBind(&createProductService); err != nil {
		msg := utils.GetValidMsg(err, &createProductService)
		ctx.JSON(http.StatusBadRequest, format.RespErrorWithData(errors.New(msg)))
	} else {
		resp, respErr := createProductService.Create(ctx.Request.Context(), files)
		if respErr != nil {
			ctx.JSON(http.StatusInternalServerError, format.RespErrorWithData(respErr))
		} else {
			ctx.JSON(http.StatusOK, format.RespSuccessWithData(resp))
		}
	}
}

// ProductListHandler 商品列表
func ProductListHandler(ctx *gin.Context) {
	listProductService := services.ProductService{}
	if err := ctx.ShouldBind(&listProductService); err != nil {
		msg := utils.GetValidMsg(err, &listProductService)
		ctx.JSON(http.StatusBadRequest, format.RespErrorWithData(errors.New(msg)))
	} else {
		resp, total, respErr := listProductService.List(ctx.Request.Context())
		if respErr != nil {
			ctx.JSON(http.StatusInternalServerError, format.RespErrorWithData(respErr))
		} else {
			ctx.JSON(http.StatusOK, format.RespListWithData(resp, total))
		}
	}
}

func ProductUpdateHandler(ctx *gin.Context) {
	
	productId := ctx.Param("productId")
	updateProductService := services.ProductService{}
	if err := ctx.ShouldBind(&updateProductService); err != nil {
		msg := utils.GetValidMsg(err, &updateProductService)
		ctx.JSON(http.StatusBadRequest, format.RespErrorWithData(errors.New(msg)))
	} else {
		resp, respErr := updateProductService.Update(ctx.Request.Context(), productId)
		if respErr != nil {
			ctx.JSON(http.StatusInternalServerError, format.RespErrorWithData(respErr))
		} else {
			ctx.JSON(http.StatusOK, format.RespSuccessWithData(resp))
		}
	}
}
