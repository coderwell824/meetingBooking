package dao

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"meetingBooking/repository/db/model"
)

type UserDao struct {
	*gorm.DB
}

// NewUserDao 创建一个用户的Dto
func NewUserDao(ctx context.Context) *UserDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &UserDao{NewDBClient(ctx)}
}

// FindUserByUserId 根据用户id找到用户
func (dao *UserDao) FindUserByUserId(id uint) (user *model.User, err error) {
	err = dao.DB.Model(&model.User{}).Where("user_id=?", id).
		First(&user).Error
	return
}

// FindUserByUsername 根据用户名找到用户
func (dao *UserDao) FindUserByUsername(userName string) (user *model.User, err error) {
	err = dao.DB.Model(&model.User{}).Where("username=?", userName).First(&user).Error
	return
}

// CreateUser 创建用户
func (dao *UserDao) CreateUser(user *model.User) (err error) {
	err = dao.DB.Model(&model.User{}).Create(user).Error
	return
}

// VerifyAdmin 验证是否为管理员
func (dao *UserDao) VerifyAdmin(user *model.User) bool {
	err := dao.DB.Model(&model.User{}).Where("is_admin=?", user.IsAdmin).Find(&user).Error
	fmt.Println(err, "user")
	return err == nil
}

// UserList 用户列表
func (dao *UserDao) UserList(start, limit int) (r []*model.User, total int64, err error) {
	err = dao.DB.Model(&model.User{}).
		Count(&total).
		Limit(limit).
		Offset((start - 1) * limit).
		Find(&r).Error
	return
}

// UpdatePassword 更新密码
func (dao *UserDao) UpdatePassword(userId uint, newPassword string) (err error) {
	err = dao.DB.Model(&model.User{}).Where("user_id", userId).Update("password", newPassword).Error
	return
}

func (dao *UserDao) DeleteUser(userId uint) (err error) {
	fmt.Println("DeleteUser", userId)
	err = dao.DB.Delete(&model.User{}, userId).Error
	return
}
