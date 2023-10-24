package reqValidator

type ReqUploadFile struct {
	File string `form:"file" binding:"required" msg:"角色名称不能为空"`
}

type ReqAddPosition struct {
	Name      string  `json:"name" binding:"required" msg:"名称不能为空"`
	Longitude float64 `json:"longitude" binding:"required" msg:"经度不能为空"`
	Latitude  float64 `json:"latitude" binding:"required" msg:"纬度不能为空"`
}

type ReqSearchPos struct {
	Longitude float64 `form:"longitude" binding:"required" msg:"经度不能为空"`
	Latitude  float64 `form:"latitude" binding:"required" msg:"纬度不能为空"`
	Radius    float64 `form:"radius" binding:"required" msg:"纬度不能为空"`
}

type ReqSendMessage struct {
	UId   string `form:"uId" binding:"required" msg:"用户id不能为空"`
	ToUid string `form:"toUid" binding:"required" msg:"接受者Id不能为空"`
	//Content string `json:"content" binding:"required" msg:"内容不能为空"`
}
