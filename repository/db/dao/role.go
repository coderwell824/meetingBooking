package dao

import (
	"context"
	"gorm.io/gorm"
	"meetingBooking/repository/db/model"
)

type RoleDao struct {
	*gorm.DB
}

// NewRoleDao 创建一个角色的Dto
func NewRoleDao(ctx context.Context) *RoleDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &RoleDao{NewDBClient(ctx)}
}

// CreateRole 创建角色
func (dao *RoleDao) CreateRole(role *model.Role) (err error) {
	err = dao.DB.Model(&model.Role{}).Create(role).Error
	return
}
