package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"meetingBooking/pkg/format"
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
