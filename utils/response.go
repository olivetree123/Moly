package utils

import (
	"bytes"
	"encoding/json"
)

// Response Moly API 返回格式
type Response struct {
	Code    int
	Message string
	Data    interface{}
}

// ToBytes response 转 []byte
func (resp *Response) ToBytes() []byte {
	content, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}
	content = bytes.ToLower(content)
	return content
}

// NewResponse 构造 Response
func NewResponse(code int, data interface{}) *Response {
	message := ""
	if code == 1000 {
		message = "BAD REQUEST"
	} else if code == 1001 {
		message = "存储数据失败"
	}
	resp := &Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
	return resp
}
