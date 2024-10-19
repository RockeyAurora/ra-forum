package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 定义API响应的JSON结构
/*
{
	"code": 10000, // 程序中的错误码
	"msg": xx,     // 提示信息
	"data": {},    // 数据
}
*/

// ResponseData 定义API响应的数据结构
type ResponseData struct {
	Code ResCode     `json:"code"` // 响应状态码
	Msg  interface{} `json:"msg"`  // 响应消息
	Data interface{} `json:"data"` // 响应数据
}

// ResponseError 返回错误响应
// 参数:
//   - c: gin的上下文
//   - code: 错误码
func ResponseError(c *gin.Context, code ResCode) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	})
}

// ResponseErrorWithMsg 返回带自定义消息的错误响应
// 参数:
//   - c: gin的上下文
//   - code: 错误码
//   - msg: 自定义错误消息
func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

// ResponseSuccess 返回成功响应
// 参数:
//   - c: gin的上下文
//   - data: 响应数据
func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	})
}
