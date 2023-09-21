package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"meetingBooking/pkg/format"
	"meetingBooking/reqValidator"
	"meetingBooking/services"
	"meetingBooking/utils"
	"net/http"
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
