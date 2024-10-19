package mysql

import "errors"

// 定义数据库操作相关的错误
var (
	// ErrorUserExist 表示用户已存在的错误
	ErrorUserExist = errors.New("用户已存在")
	
	// ErrorUserNotExist 表示用户不存在的错误
	ErrorUserNotExist = errors.New("用户不存在")
	
	// ErrorInvalidPassword 表示用户名或密码错误的错误
	ErrorInvalidPassword = errors.New("用户名或密码错误")
	
	// ErrorInvalidID 表示无效ID的错误
	ErrorInvalidID = errors.New("无效的ID")
)
