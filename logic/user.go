package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/jwt"
	"bluebell/pkg/rabbitmq"
	"bluebell/pkg/snowflake"
	"fmt"
	"go.uber.org/zap"
)

// 用户相关的业务逻辑代码

// SignUp 处理用户注册逻辑
// 参数:
//   - p: 用户注册参数
//
// 返回值:
//   - error: 可能发生的错误
func SignUp(p *models.ParamSignUp) (err error) {
	// 1. 检查用户是否已存在
	if err := mysql.CheckUserExist(p.Username); err != nil {
		return err
	}
	// 2. 生成用户唯一ID
	userID := snowflake.GenID()
	// 3. 构造User实例
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	// 4. 将用户信息保存到数据库
	return mysql.InsertUser(user)
}

// Login 处理用户登录逻辑
// 参数:
//   - p: 用户登录参数
//
// 返回值:
//   - *models.User: 登录成功的用户信息
//   - error: 可能发生的错误
func Login(p *models.ParamLogin) (user *models.User, err error) {
	user = &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	// 验证用户登录信息
	if err := mysql.Login(user); err != nil {
		return nil, err
	}
	// 生成JWT令牌
	token, err := jwt.GenToken(user.UserID, user.Username)
	if err == nil {
		user.Token = token
	}
	return
}

func SignUpNew(p *models.ParamSignUp) (err error) {
	var errs error

	if p.Email != "" {
		Ed := &models.ParamEmailData{
			Email:    p.Email,
			Username: p.Username,
			Password: p.Password,
		}
		zap.L().Debug("emaildetail", zap.String("Username", Ed.Username), zap.String("Email", Ed.Email))
		// 使用生产者发布邮件任务到队列
		err = rabbitmq.PublishEmailTask(Ed)
		if err != nil {
			errs = fmt.Errorf("failed to publish email task: %w; %v", err, errs)
		}
	}

	if err = SignUp(p); err != nil {
		errs = fmt.Errorf("signup error: %w; %v", err, errs)
	}
	zap.L().Debug("signup success", zap.String("email", p.Email), zap.String("username", p.Username))
	return errs
}
