package services

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"meetingBooking/pkg/format"
	"meetingBooking/repository/db/dao"
	"meetingBooking/repository/db/model"
	"meetingBooking/reqValidator"
	"meetingBooking/utils"
)

func UserRegister(ctx context.Context, req *reqValidator.ReqRegister) (response interface{}, err error) {
	err = utils.ValidEmail(req.Email, req.Captcha)
	if err != nil {
		return
	}
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.FindUserByUsername(req.Username)
	switch err {
	case gorm.ErrRecordNotFound:
		user = &model.User{
			Username: req.Username,
		}
		if err = user.SetPassword(req.Password); err != nil {
			return
		}
		if err = userDao.CreateUser(user); err != nil {
			return
		}
		return format.RespSuccessWithData("注册成功"), nil
	case nil:
		err = errors.New("用户已存在")
		log.Println(err)
		return
	default:
		return
	}
}
func UserLogin(ctx context.Context, req *reqValidator.ReqLogin) (response interface{}, err error) {
	
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.FindUserByUsername(req.Username)
	if err != nil {
		return format.RespSuccessWithData("用户未注册"), nil
	}
	if isCorrect := user.CheckPassword(req.Password); isCorrect != true {
		return format.RespSuccessWithData("密码错误"), nil
	}
	token := utils.GenerateToken(user.ID, req.Username, 0)
	if err != nil {
		log.Println("Error generating access token")
	}
	fmt.Println(token)
	return format.RespSuccessWithData(token), nil
}
