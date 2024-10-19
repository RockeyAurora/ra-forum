package controller

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CtxUserIDKey 是用于在上下文中存储用户ID的键
const CtxUserIDKey = "userID"

// ErrorUserNotLogin 是用户未登录时返回的错误
var ErrorUserNotLogin = errors.New("用户未登录")

// getCurrentUserID 获取当前登录的用户ID
// 参数:
//   - c: gin的上下文
// 返回值:
//   - userID: 用户ID
//   - err: 可能的错误，如用户未登录
func getCurrentUserID(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(CtxUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}

// GetPageInfo 从请求中获取分页信息
// 参数:
//   - c: gin的上下文
// 返回值:
//   - page: 页码，默认为1
//   - size: 每页大小，默认为10
func GetPageInfo(c *gin.Context) (int64, int64) {
	PageStr := c.Query("page")
	SizeStr := c.Query("size")
	var (
		page int64
		size int64
		err  error
	)
	page, err = strconv.ParseInt(PageStr, 10, 64)
	if err != nil {
		page = 1 // 如果解析失败，默认为第1页
	}
	size, err = strconv.ParseInt(SizeStr, 10, 64)
	if err != nil {
		size = 10 // 如果解析失败，默认每页10条
	}
	return page, size
}
