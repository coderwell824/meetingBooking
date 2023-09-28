package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"meetingBooking/pkg/format"
	"meetingBooking/repository/cache"
	"meetingBooking/reqValidator"
	"meetingBooking/services"
	"meetingBooking/utils"
	"net/http"
	"strconv"
	"time"
)

func UserRegisterHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req reqValidator.ReqRegister
		if err := ctx.ShouldBind(&req); err == nil {
			resp, respErr := services.UserRegister(ctx.Request.Context(), &req)
			if respErr != nil {
				ctx.JSON(http.StatusInternalServerError, format.RespErrorWithData(respErr))
			} else {
				ctx.JSON(http.StatusOK, resp)
			}
		} else {
			msg := utils.GetValidMsg(err, &req)
			ctx.JSON(http.StatusBadRequest, format.RespErrorWithData(errors.New(msg)))
		}
	}
}

func SendEmailHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email := ctx.Query("email")
		if email == "" {
			ctx.JSON(http.StatusBadRequest, format.RespErrorWithData(errors.New("邮箱不能为空")))
		} else {
			//产生六位数验证码
			rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
			emailCode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
			err := cache.RedisSetKey(fmt.Sprintf("captcha_%s", email), emailCode, 600000*time.Second)
			if err != nil {
				log.Panicln("Redis设置key失败")
			}
			err = services.SendEmailService(emailCode, []string{email})
			if err != nil {
				log.Panicln("验证码发送失败")
			}
			ctx.JSON(http.StatusOK, format.RespSuccessWithData("验证码已发送"))
		}
	}
}

func UserLoginHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req reqValidator.ReqLogin
		if err := ctx.ShouldBind(&req); err != nil {
			msg := utils.GetValidMsg(err, &req)
			ctx.JSON(http.StatusBadRequest, format.RespErrorWithData(errors.New(msg)))
		} else {
			resp, respErr := services.UserLogin(ctx.Request.Context(), &req)
			if respErr != nil {
				ctx.JSON(http.StatusInternalServerError, format.RespErrorWithData(respErr))
			} else {
				ctx.JSON(http.StatusOK, format.RespSuccessWithData(resp))
			}
		}
	}
}

func GetUserListHandler() gin.HandlerFunc {
	var req reqValidator.ReqUserList
	return func(ctx *gin.Context) {
		if err := ctx.ShouldBindQuery(&req); err != nil {
			msg := utils.GetValidMsg(err, &req)
			ctx.JSON(http.StatusBadRequest, format.RespErrorWithData(errors.New(msg)))
		} else {
			resp, total, respErr := services.UserList(ctx.Request.Context(), &req)
			if respErr != nil {
				ctx.JSON(http.StatusInternalServerError, format.RespErrorWithData(respErr))
			} else {
				ctx.JSON(http.StatusOK, format.RespListWithData(resp, total))
			}
		}
	}
}

func UpdatePasswordHandler() gin.HandlerFunc {
	var req reqValidator.ReqUpdatePassword
	return func(ctx *gin.Context) {
		userId := ctx.Param("userId")
		if userId == "" {
			ctx.JSON(http.StatusInternalServerError, format.RespErrorWithData(errors.New("用户id不存在")))
		}
		newUserId, _ := strconv.Atoi(userId)
		if err := ctx.ShouldBind(&req); err != nil {
			msg := utils.GetValidMsg(err, &req)
			ctx.JSON(http.StatusBadRequest, format.RespErrorWithData(errors.New(msg)))
		} else {
			respErr := services.UserUpdatePassword(ctx.Request.Context(), &req, uint(newUserId))
			if respErr != nil {
				ctx.JSON(http.StatusInternalServerError, format.RespErrorWithData(respErr))
			} else {
				ctx.JSON(http.StatusOK, format.RespSuccessWithData("密码更新成功"))
			}
		}
	}
}

func DeleteUserHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.Param("userId")
		if userId == "" {
			ctx.JSON(http.StatusInternalServerError, format.RespErrorWithData(errors.New("用户id不存在")))
		}
		newUserId, _ := strconv.Atoi(userId)
		_, respErr := services.DeleteUserById(ctx.Request.Context(), uint(newUserId))
		if respErr != nil {
			ctx.JSON(http.StatusInternalServerError, format.RespErrorWithData(respErr))
		} else {
			ctx.JSON(http.StatusOK, format.RespSuccessWithData("用户删除成功"))
		}
	}
}

func UpdateUserInfoHandler() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var req reqValidator.ReqUpdateInfo
		if err := ctx.ShouldBind(&req); err != nil {
			msg := utils.GetValidMsg(err, &req)
			ctx.JSON(http.StatusBadRequest, format.RespErrorWithData(errors.New(msg)))
		} else {
			respErr := services.UserUpdatePassword(ctx.Request.Context(), &req)
			if respErr != nil {
				ctx.JSON(http.StatusInternalServerError, format.RespErrorWithData(respErr))
			} else {
				ctx.JSON(http.StatusOK, format.RespSuccessWithData("密码更新成功"))
			}
		}
	}
}
