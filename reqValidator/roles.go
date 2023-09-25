package reqValidator

type CreateRoleReq struct {
	RoleName string `json:"roleName" binding:"required" msg:"角色名称不能为空"`
}
