package services

import (
	"context"
	"errors"
	"log"
	"meetingBooking/repository/db/dao"
	"meetingBooking/repository/db/model"
	"meetingBooking/reqValidator"
	"meetingBooking/utils"
)

func CreateRole(ctx context.Context, req *reqValidator.CreateRoleReq) (response interface{}, err error) {
	
	u, err := utils.GetUserInfo(ctx)
	if err != nil {
		log.Println(err)
		return
	}
	roleDao := dao.NewRoleDao(ctx)
	user, err := dao.NewUserDao(ctx).FindUserByUserId(u.Id)
	if user.IsAdmin == false {
		err = errors.New("user is not a admin")
		return
	}
	role := &model.Role{
		Name: req.RoleName,
		Users: []model.User{
			{ID: 1},
		},
	}
	if err = roleDao.CreateRole(role); err != nil {
		log.Println(err, "create")
	}
	
	return "角色创建成功", nil
}
