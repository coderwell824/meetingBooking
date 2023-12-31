package reqValidator

type ReqRegister struct {
	Username string `json:"username" binding:"required" msg:"用户名不能为空"`
	NickName string `json:"nickname" binding:"required" msg:"昵称不能为空"`
	Password string `json:"password" binding:"required,min=6" msg:"密码长度至少6位"`
	Email    string `json:"email" binding:"required" msg:"邮箱不能为空"` //TODO： 邮箱加验证
	Captcha  string `json:"captcha" binding:"required" msg:"验证码不能为空"`
}

type ReqLogin struct {
	Username string `json:"username" binding:"required" msg:"用户名不能为空"`
	Password string `json:"password" binding:"required,min=6" msg:"密码长度至少6位"`
}

type ReqUserList struct {
	PageNum  int `form:"pageNum" binding:"required" msg:"pageNum不能为空"` //Query参数使用的是Form
	PageSize int `form:"pageSize" binding:"required" msg:"pageSize不能为空"`
}
type ReqUpdatePassword struct {
	Password    string `json:"password" binding:"required,min=6" msg:"密码长度至少6位"`
	NewPassword string `json:"newPassword" binding:"required,min=6" msg:"新密码长度至少6位"`
}
type ReqUpdateInfo struct {
	Nickname    string `json:"nickname" binding:"required" msg:"昵称不能为空"`
	Email       string `json:"email" binding:"required" msg:"邮箱不能为空"`
	AvatarUrl   string `json:"avatarUrl" binding:"required" msg:"头像不能为空"`
	PhoneNumber string `json:"phoneNumber" binding:"required" msg:"手机号不能为空"`
}
