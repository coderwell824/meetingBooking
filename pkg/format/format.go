package format

import "net/http"

type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Error   string      `json:"error"`
	Total   int64       `json:"total"`
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

// RespListWithData   带List成功返回
func RespListWithData(data interface{}, total int64) *Response {
	r := &Response{
		Code:    http.StatusOK,
		Data:    data,
		Total:   total,
		Message: "查询成功",
	}
	return r
}

// RespErrorWithData   带error错误返回
func RespErrorWithData(err error) *Response {
	r := &Response{
		Code:  http.StatusOK,
		Error: err.Error(),
	}
	return r
}
