package services

import (
	"context"
	"fmt"
	"meetingBooking/repository/db/dao"
	"meetingBooking/reqValidator"
	"meetingBooking/utils"
)

func UserRegister(ctx context.Context, req *reqValidator.ReqRegister) (interface{}, error) {
	err := utils.ValidEmail(req.Email, req.Captcha)
	if err != nil {
		return "", err
	}
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.FindUserByUsername(req.Username)
	
	fmt.Println(user, err.Error(), "ddsdsdsdsd")
	
	return nil, nil
	
}
