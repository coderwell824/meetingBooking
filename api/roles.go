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

func CreateRoleHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req reqValidator.CreateRoleReq
		if err := ctx.ShouldBind(&req); err != nil {
			msg := utils.GetValidMsg(err, &req)
			ctx.JSON(http.StatusBadRequest, format.RespErrorWithData(errors.New(msg)))
		} else {
			resp, respErr := services.CreateRole(ctx.Request.Context(), &req)
			if respErr != nil {
				ctx.JSON(http.StatusInternalServerError, format.RespErrorWithData(respErr))
			} else {
				ctx.JSON(http.StatusOK, resp)
			}
		}
	}
}
