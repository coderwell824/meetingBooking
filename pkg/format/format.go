package format

import "net/http"

type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Error   string      `json:"error"`
}

// RespSuccessWithData   带data成功返回
func RespSuccessWithData(data interface{}) *Response {
	
	r := &Response{
		Code:    http.StatusOK,
		Data:    data,
		Message: "操作成功",
	}
	
	return r
}

// RespErrorWithData   带data成功返回
func RespErrorWithData(err error) *Response {
	r := &Response{
		Code:  http.StatusOK,
		Data:  "",
		Error: err.Error(),
	}
	return r
}
