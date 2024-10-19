package mysql

import (
	"bluebell/setting"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// db 是全局的数据库连接对象
var db *sqlx.DB

// Init 初始化MySQL连接
// 参数:
//   - cfg: MySQL配置信息
// 返回值:
//   - err: 可能的错误
func Init(cfg *setting.MySQLConfig) (err error) {
	// 构建数据源名称（DSN）
	// 格式: "user:password@tcp(host:port)/dbname?parseTime=true&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	
	// 连接到MySQL数据库
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}
	
	// 设置数据库连接池的最大打开连接数
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	
	// 设置数据库连接池的最大空闲连接数
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	
	return
}

// Close 关闭MySQL连接
// 注意: 这个函数应该在程序退出时调用
func Close() {
	_ = db.Close()
}
