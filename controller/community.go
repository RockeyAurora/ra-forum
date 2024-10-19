package controller

import (
	"bluebell/logic"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ---- 社区相关的处理函数 ----

// CommunityHandler 处理获取社区列表的请求
func CommunityHandler(c *gin.Context) {
	// 查询所有社区（community_id, community_name），以列表形式返回
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("获取社区列表失败", zap.Error(err))
		ResponseError(c, CodeServerBusy) // 不向客户端暴露具体错误信息
		return
	}
	ResponseSuccess(c, data)
}

// CommunityDetailHandler 处理获取社区详情的请求
func CommunityDetailHandler(c *gin.Context) {
	// 1. 从URL参数中获取社区ID
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 2. 根据ID获取社区详情
	data, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("获取社区详情失败", zap.Error(err))
		ResponseError(c, CodeServerBusy) // 不向客户端暴露具体错误信息
		return
	}
	ResponseSuccess(c, data)
}
