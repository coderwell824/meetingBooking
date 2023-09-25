package services

import (
	"context"
	"fmt"
	"log"
	"meetingBooking/repository/db/dao"
	"meetingBooking/reqValidator"
	"meetingBooking/utils"
)

func CreateRole(ctx context.Context, req *reqValidator.CreateRoleReq) (response interface{}, err error) {
	
	u, err := utils.GetUserInfo(ctx)
	if err != nil {
		log.Println(err)
		return
	}
	user, err := dao.NewUserDao(ctx).FindUserByUserId(u.Id)
	fmt.Println(user, "xxxxs")
	//dao.UserDao.VerifyAdmin(user)
	return
}
