package model

import "time"

// Role 角色表
type Role struct {
	ID          uint         `json:"roleId" gorm:"column:role_id"` //角色ID
	Name        string       `json:"name" gorm:"type:varchar(20)"` //角色名称
	Users       []User       `json:"users"`
	Permissions []Permission `gorm:"many2many:role_permission;joinForeignKey:RoleID;JoinReferences:PermissionID"`
	CreatedAt   time.Time    `json:"createTime"`
	UpdatedAt   time.Time    `json:"updateTime"`
}

// Permission 权限表
type Permission struct {
	ID          uint      `json:"permissionId" gorm:"column:permission_id"` //权限ID
	Code        string    `json:"code" gorm:"type:varchar(20)"`             //权限代码
	Description string    `json:"description" gorm:"type:varchar(100)"`     //描述
	Roles       []Role    `gorm:"many2many:role_permissions;joinForeignKey:PermissionID;JoinReferences:RoleID"`
	CreatedAt   time.Time `json:"createTime"`
	UpdatedAt   time.Time `json:"updateTime"`
}

// RoleUser 用户角色中间表
type RoleUser struct {
	ID        uint      `json:"UserRoleId" gorm:"column:user_role_id"` //ID
	UserID    uint      `json:"userId" gorm:"primaryKey"`
	RoleID    uint      `json:"roleId" gorm:"primaryKey"`
	CreatedAt time.Time `json:"createTime"`
	UpdatedAt time.Time `json:"updateTime"`
}

// RolePermission 角色权限中间表
type RolePermission struct {
	ID           uint      `json:"rolePermissionId" gorm:"column:role_permission_id"` //ID
	RoleID       uint      `json:"roleId" gorm:"primaryKey"`
	PermissionID uint      `json:"permissionId" gorm:"primaryKey"`
	CreatedAt    time.Time `json:"createTime"`
	UpdatedAt    time.Time `json:"updateTime"`
}
