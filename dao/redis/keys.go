package redis

// Redis键定义
// 注意：使用命名空间方式定义Redis键，以便于查询和拆分

// 常量定义
const (
	// Prefix 项目前缀
	Prefix = "bluebell:"

	// KeyPostTimeZSet 帖子时间有序集合
	// 类型：zset
	// 用途：存储帖子ID及其发布时间
	KeyPostTimeZSet = "post:time"

	// KeyPostScoreZSet 帖子分数有序集合
	// 类型：zset
	// 用途：存储帖子ID及其投票分数
	KeyPostScoreZSet = "post:score"

	// KeyPostVotedPF 用户投票记录
	// 类型：zset
	// 用途：记录用户对帖子的投票类型
	KeyPostVotedPF = "post:voted:"

	KeyCommunitySetPF = "community:" // 类型：set  保存每个分区下帖子id
)

// getRedisKey 获取完整的Redis键
// 参数：
//   - key: 键名
//
// 返回值：
//   - 带有项目前缀的完整键名
func getRedisKey(key string) string {
	return Prefix + key
}
