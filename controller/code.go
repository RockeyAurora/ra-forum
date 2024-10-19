package controller

// ResCode 定义响应状态码类型
type ResCode int64

// 定义各种响应状态码常量
const (
	CodeSuccess ResCode = 1000 + iota // 成功
	CodeInvalidParam                  // 无效参数
	CodeUserExist                     // 用户已存在
	CodeUserNotExist                  // 用户不存在
	CodeInvalidPassword               // 无效密码
	CodeServerBusy                    // 服务器繁忙

	CodeNeedLogin    // 需要登录
	CodeInvalidToken // 无效令牌
)

// 定义状态码对应的消息映射
var codeMsgMap = map[ResCode]string{
	CodeSuccess:         "success",
	CodeInvalidParam:    "请求参数错误",
	CodeUserExist:       "用户名已存在",
	CodeUserNotExist:    "用户名不存在",
	CodeInvalidPassword: "用户名或密码错误",
	CodeServerBusy:      "服务繁忙",

	CodeNeedLogin:    "需要登录",
	CodeInvalidToken: "无效的token",
}

// Msg 方法返回对应状态码的消息
func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
