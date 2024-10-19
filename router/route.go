package router

import (
	"bluebell/controller"
	"bluebell/logger"
	"bluebell/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true)) // 使用自定义的日志和恢复中间件

	v1 := r.Group("/api/v1") // 创建API v1版本的路由组

	// 无需认证的路由
	v1.POST("/signup", controller.SignUpHandler) // 用户注册
	v1.POST("/login", controller.LoginHandler)   // 用户登录

	v1.Use(middlewares.JWTAuthMiddleware()) // 应用JWT认证中间件

	{
		// 社区相关路由
		v1.GET("/community", controller.CommunityHandler)           // 获取社区列表
		v1.GET("/community/:id", controller.CommunityDetailHandler) // 获取特定社区详情

		// 帖子相关路由
		v1.POST("/post", controller.CreatePostHandler)       // 创建新帖子
		v1.GET("/post/:id", controller.GetPostDetailHandler) // 获取特定帖子详情
		v1.GET("/posts/", controller.GetPostListHandler)     // 获取帖子列表
		v1.GET("/posts2/", controller.GetPostListHandler2)   // 根据时间或分数获取帖子列表（优化版）
		v1.POST("/vote", controller.PostVoteController)      // 帖子投票
	}

	// 处理404路由
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
