package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"meetingBooking/pkg/format"
	"meetingBooking/utils"
	"net/http"
	"strings"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var code int = 200
		token := ctx.GetHeader("Authorization")
		if token == "" {
			code = http.StatusNotFound
			ctx.JSON(http.StatusBadRequest, format.RespErrorWithData(errors.New("缺少token")))
			ctx.Abort()
			return
		}
		token = strings.Split(token, " ")[1]
		claims, err := utils.ParseToken(token)
		if err != nil {
			code = 402
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = 407
		}
		if code != 200 {
			ctx.JSON(http.StatusOK, format.RespErrorWithData(errors.New("token已过期")))
			ctx.Abort()
			return
		}
		ctx.Request = ctx.Request.WithContext(utils.NewContext(ctx.Request.Context(), &utils.UserInfo{Id: claims.Id}))
		ctx.Next()
	}
}
