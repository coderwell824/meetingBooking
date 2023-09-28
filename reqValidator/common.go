package reqValidator

type ReqUploadFile struct {
	File string `form:"file" binding:"required" msg:"角色名称不能为空"`
}
