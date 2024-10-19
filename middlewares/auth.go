package middlewares

import (
	"bluebell/controller"
	"bluebell/pkg/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式：
		// 1. 放在请求头（Header）
		// 2. 放在请求体（Body）
		// 3. 放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 例如：Authorization: Bearer xxxxxxx.xxx.xxx 或 X-TOKEN: xxx.xxx.xx
		// 注意：具体实现方式应根据实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			// 如果请求头中没有Authorization字段，则返回需要登录的错误
			controller.ResponseError(c, controller.CodeNeedLogin)
			c.Abort()
			return
		}

		// 按空格分割Authorization字段的值
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			// 如果Authorization格式不正确，则返回无效的token错误
			controller.ResponseError(c, controller.CodeInvalidToken)
			c.Abort()
			return
		}

		// parts[1]是获取到的tokenString，使用之前定义好的ParseToken函数来解析它
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			// 如果token解析失败，则返回无效的token错误
			controller.ResponseError(c, controller.CodeInvalidToken)
			c.Abort()
			return
		}

		// 将当前请求的userID信息保存到请求的上下文c中
		c.Set(controller.CtxUserIDKey, mc.UserID)

		// 继续处理请求
		c.Next()
		// 注意：在后续的处理请求的函数中，可以通过c.Get(CtxUserIDKey)来获取当前请求的用户信息
	}
}
