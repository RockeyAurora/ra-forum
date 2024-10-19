// Package snowflake 提供了基于Twitter的Snowflake算法的唯一ID生成功能
package snowflake

import (
	"time"

	sf "github.com/bwmarrin/snowflake"
)

// node 是用于生成ID的Snowflake节点
var node *sf.Node

// Init 初始化Snowflake节点
// 参数:
//   - startTime: 起始时间，格式为"2006-01-02"
//   - machineID: 机器ID
// 返回:
//   - error: 可能发生的错误
func Init(startTime string, machineID int64) (err error) {
	var st time.Time
	// 解析起始时间
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}
	// 设置Snowflake的纪元（起始时间）
	sf.Epoch = st.UnixNano() / 1000000
	// 创建一个新的Snowflake节点
	node, err = sf.NewNode(machineID)
	return
}

// GenID 生成一个新的唯一ID
// 返回:
//   - int64: 生成的唯一ID
func GenID() int64 {
	return node.Generate().Int64()
}
