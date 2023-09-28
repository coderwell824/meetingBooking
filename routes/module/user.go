package module

import (
	"github.com/gin-gonic/gin"
	"meetingBooking/api"
)

func LoadLoginRoute(v1 *gin.RouterGroup) {

	user := v1.Group("/user")

	user.POST("register", api.UserRegisterHandler())
	user.POST("login", api.UserLoginHandler())
	user.GET("sendEmail", api.SendEmailHandler())

}

func LoadUserRoute(v1 *gin.RouterGroup) {
	user := v1.Group("/user")
	user.POST("info/update",api.UpdateUserInfoHandler()) )
	user.PUT("updatePassword/:userId", api.UpdatePasswordHandler())
	user.GET("list", api.GetUserListHandler())
	user.GET("freeze", func(context *gin.Context) {
		context.JSON(200, "success")
	})
	user.DELETE("delete/:userId", api.DeleteUserHandler())
}
