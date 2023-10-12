package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"meetingBooking/pkg/format"
	"meetingBooking/repository/cache"
	"meetingBooking/reqValidator"
	"meetingBooking/utils"
	"net/http"
	"os"
)

func UploadFileHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		file, err := ctx.FormFile("file")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, format.RespErrorWithData(errors.New("文件上传失败")))
		}
		err = ctx.SaveUploadedFile(file, "./assets/images/"+file.Filename) //把用户上传的文件存到data目录下
		if err != nil {
			ctx.JSON(http.StatusOK, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, format.RespSuccessWithData("http://"+ctx.Request.Host+"/assets/images/"+file.Filename))
	}
}

func GetImageHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		filename := ctx.Param("filename")
		imagePath := "assets/images/" + filename
		
		// 检查文件是否存在
		_, err := os.Stat(imagePath)
		if os.IsNotExist(err) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "文件不存在"})
			return
		}
		
		// 设置响应头，告诉浏览器返回的是图片
		ctx.Header("Content-Type", "image/jpeg") // 这里假设上传的图片是JPEG格式
		
		// 读取图片文件并将其写入响应体
		ctx.File(imagePath)
	}
}
func AddPosHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req reqValidator.ReqAddPosition
		if err := ctx.ShouldBind(&req); err != nil {
			msg := utils.GetValidMsg(err, &req)
			ctx.JSON(http.StatusBadRequest, format.RespErrorWithData(errors.New(msg)))
		} else {
			err := cache.RedisGeoAdd(&redis.GeoLocation{
				Name:      req.Name,
				Longitude: req.Longitude,
				Latitude:  req.Latitude,
			})
			if err != nil {
				ctx.JSON(http.StatusBadGateway, format.RespErrorWithData(errors.New("位置信息添加失败")))
			}
			ctx.JSON(http.StatusOK, format.RespSuccessWithData("位置信息添加成功"))
			
		}
	}
}

func GetPosHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		name := ctx.Param("name")
		if name == "" {
			ctx.JSON(http.StatusBadRequest, format.RespErrorWithData(errors.New("缺少name参数")))
		}
		
		resPos, err := cache.RedisGetGeo(name)
		if err != nil {
			ctx.JSON(http.StatusBadGateway, format.RespErrorWithData(errors.New("位置信息获取失败")))
		}
		for _, pos := range resPos {
			ctx.JSON(http.StatusOK, format.RespSuccessWithData(pos))
		}
	}
}

func GetPosAllHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
	
	}
}

func GetNearbySearchHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req reqValidator.ReqSearchPos
		if err := ctx.ShouldBindQuery(&req); err != nil {
			msg := utils.GetValidMsg(err, &req)
			ctx.JSON(http.StatusBadRequest, format.RespErrorWithData(errors.New(msg)))
		} else {
			resRadius, err := cache.RedisGeoRadius(req.Longitude, req.Latitude, req.Radius)
			if err != nil {
				ctx.JSON(http.StatusBadGateway, format.RespErrorWithData(errors.New("范围信息查询失败")))
			}
			ctx.JSON(http.StatusOK, format.RespSuccessWithData(resRadius))
			
		}
	}
}
