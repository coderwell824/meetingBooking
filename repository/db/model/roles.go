package model

import "time"

//Role 角色表
type Role struct {
	ID         uint       `json:"roleId" gorm:"column:role_id"` //角色ID
	Name       string     `json:"name" gorm:"type:varchar(20)"` //角色名称
	CreateTime *time.Time `json:"createTime"`
	UpdateTime *time.Time `json:"updateTime"`
	//Users      []User
}

//Permission 权限表
type Permission struct {
	ID          uint       `json:"permissionId" gorm:"column:permission_id"` //权限ID
	Code        string     `json:"code" gorm:"type:varchar(20)"`             //权限代码
	Description string     `json:"description" gorm:"type:varchar(100)"`     //描述
	CreateTime  *time.Time `json:"createTime"`
	UpdateTime  *time.Time `json:"updateTime"`
}

//UserRole 用户角色中间表
type UserRole struct {
	ID     uint `json:"UserRoleId" gorm:"column:user_role_id"` //ID
	UserID uint `json:"userId"`
	RoleID uint `json:"roleId"`
}

//RolePermission 角色权限中间表
type RolePermission struct {
	ID           uint `json:"rolePermissionId" gorm:"column:role_permission_id"` //ID
	RoleID       uint `json:"roleId"`
	PermissionID uint `json:"permissionId"`
}
