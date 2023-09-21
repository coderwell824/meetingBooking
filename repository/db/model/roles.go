package model

//Roles 角色表
type Roles struct {
	ID   uint   `json:"roleId" gorm:"column:role_id"` //角色ID
	Name string `json:"name" gorm:"type:varchar(20)"` //角色名称
}

//Permissions 权限表
type Permissions struct {
	ID          uint   `json:"roleId" gorm:"column:permission_id"`   //权限ID
	Code        string `json:"code" gorm:"type:varchar(20)"`         //权限代码
	Description string `json:"description" gorm:"type:varchar(100)"` //描述
}

//UserRoles 用户角色中间表
type UserRoles struct {
	ID      uint `json:"UserRolesId" gorm:"column:user_roles_id"` //ID
	UserID  uint `json:"userId"`
	RolesID uint `json:"rolesId"`
}

//RolePermissions 角色权限中间表
type RolePermissions struct {
	ID            uint `json:"UserRolesId" gorm:"column:role_permission_id"` //ID
	RolesID       uint `json:"rolesId"`
	PermissionsID uint `json:"permissionsId"`
}
