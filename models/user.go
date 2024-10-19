// User 结构体定义了用户模型
package models

type User struct {
	UserID   int64  `db:"user_id"`   // 用户ID
	Username string `db:"username"`  // 用户名
	Password string `db:"password"`  // 密码
	Token    string                  // 用户令牌
}
