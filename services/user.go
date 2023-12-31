package services

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"meetingBooking/config"
	"meetingBooking/consts"
	"meetingBooking/repository/db/dao"
	"meetingBooking/repository/db/model"
	"meetingBooking/reqValidator"
	"meetingBooking/utils"
	"mime/multipart"
)

type UserService struct {
}

func (u UserService) UserRegister(ctx context.Context, req *reqValidator.ReqRegister) (response interface{}, err error) {
	err = utils.ValidEmail(req.Email, req.Captcha)
	if err != nil {
		return
	}
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.FindUserByUsername(req.Username)
	switch err {
	case gorm.ErrRecordNotFound:
		
		newPassword, err := user.SetPassword(req.Password)
		if err != nil {
			log.Println("Error setting")
		}
		user = &model.User{
			Username: req.Username,
			Nickname: req.NickName,
			Email:    req.Email,
			Password: newPassword,
		}
		if err = userDao.CreateUser(user); err != nil {
			log.Println("Error setting", err)
		}
		return "注册成功", nil
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
		err = errors.New("用户未注册")
		return
	}
	// TODO: 密码修改后登录有问题
	if isCorrect := user.CheckPassword(req.Password); isCorrect != true {
		return "密码错误", nil
	}
	response = utils.GenerateToken(user.ID, req.Username, 0)
	if err != nil {
		log.Println("Error generating access token")
	}
	return
}

func UserList(ctx context.Context, req *reqValidator.ReqUserList) (resp interface{}, total int64, err error) {
	
	userDao := dao.NewUserDao(ctx)
	resp, total, err = userDao.UserList(req.PageNum, req.PageSize)
	if err != nil {
		log.Println("Error getting user list")
	}
	return
}
func UserUpdatePassword(ctx context.Context, req *reqValidator.ReqUpdatePassword, userId uint) (err error) {
	
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.FindUserByUserId(userId)
	if err != nil {
		return errors.New("用户不存在")
	}
	if isCorrect := user.CheckPassword(req.Password); isCorrect != true {
		return errors.New("密码错误")
	}
	newPassword, err := user.SetPassword(req.Password)
	if err != nil {
		return
	}
	err = userDao.UpdatePassword(userId, newPassword)
	return
}

func DeleteUserById(ctx context.Context, userId uint) (resp interface{}, err error) {
	
	u, err := utils.GetUserInfo(ctx)
	if err != nil {
		log.Println(err)
		return
	}
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.FindUserByUserId(u.Id)
	if user.IsAdmin == false {
		err = errors.New("不是管理员")
		return
	}
	err = userDao.DeleteUser(userId)
	if err != nil {
		log.Println("delete user failed")
	}
	return "删除用户成功", nil
	
}

func UpdateUserInfo(ctx context.Context, req *reqValidator.ReqUpdateInfo) (resp interface{}, err error) {
	u, err := utils.GetUserInfo(ctx)
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.FindUserByUserId(u.Id)
	fmt.Println(user, "update user")
	return
}

func UploadAvatarService(ctx context.Context, file multipart.File, fileSize int64) (resp interface{}, err error) {
	var path string
	
	u, err := utils.GetUserInfo(ctx)
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.FindUserByUserId(u.Id)
	if err != nil {
		err = errors.New("用户信息获取失败")
		return
	}
	
	//兼容两种存储方式
	if config.UploadModel == consts.UploadModelLocal {
		path, err = utils.UploadAvatarToLocalStatic(file, user.ID, user.Username)
	} else {
		path, err = utils.UploadImageToQiQiu(file, fileSize)
	}
	
	if err != nil {
		err = errors.New("图片存储失败")
		return
	}
	
	user.AvatarUrl = path
	err = userDao.UpdateUserById(user.ID, user)
	if err != nil {
		err = errors.New("图片更新失败")
		return
	}
	
	resp = "图片上传成功"
	
	return
}
