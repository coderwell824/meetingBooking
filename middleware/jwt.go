package middleware

//
//import (
//	"github.com/gin-gonic/gin"
//	"meetingBooking/pkg/ctl"
//	"meetingBooking/pkg/e"
//	"meetingBooking/utils"
//	"net/http"
//	"time"
//)
//
//func JWT() gin.HandlerFunc {
//
//	return func(ctx *gin.Context) {
//		var code int
//		code = e.SUCCESS
//		token := ctx.GetHeader("Authorization")
//
//		if token == "" {
//			code = http.StatusNotFound
//			ctx.JSON(e.InvalidParams, ctl.RespSuccessWithData("token不存在", e.AuthTokenNotFound))
//			return
//		}
//
//		claims, err := utils.ParseToken(token)
//		if err != nil {
//			code = e.ErrorAuthCheckTokenFail
//		} else if time.Now().Unix() > claims.ExpiresAt {
//			code = e.ErrorAuthCheckTokenTimeout
//		}
//
//		if code != e.SUCCESS {
//			ctx.JSON(e.InvalidParams, ctl.RespSuccessWithData("token已过期", e.ErrorAuthCheckTokenTimeout))
//			return
//		}
//
//		ctx.Request = ctx.Request.WithContext(ctl.NewContext(ctx.Request.Context(), &ctl.UserInfo{Id: claims.Id}))
//		ctx.Next()
//
//	}
//}
