package logic

import (
	"bluebell/dao/redis"
	"bluebell/models"
	"go.uber.org/zap"
	"strconv"
)

// 投票功能
// 简化版投票算法：
// - 用户投一票加432分
// - 86400/432=200，需要200张赞成票才能给帖子续一天

// 投票的几种情况：
// 1. direction = 1 (赞成票)
//    - 之前没有投过票，现在投赞成票
//    - 之前投反对票，现在改投赞成票
// 2. direction = 0 (取消投票)
//    - 之前投赞成票，现在要取消
//    - 之前投反对票，现在要取消
// 3. direction = -1 (反对票)
//    - 之前没有投过票，现在投反对票
//    - 之前投赞成票，现在改投反对票

// 投票限制：
// - 每个帖子发表之日起一个星期之内允许用户投票
// - 到期之后将Redis中保存的赞成票数及反对票数存储在MySQL，并删除KeyPostVotedPF

// VoteForPost 处理用户对帖子的投票
// 参数：
//   - userID: 用户ID
//   - p: 投票参数
// 返回值：
//   - error: 可能发生的错误
func VoteForPost(userID int64, p *models.ParamVoteData) error {
	// TODO: 判断投票限制
	// TODO: 更新分数
	// TODO: 记录用户投票数据
	
	// 记录调试日志
	zap.L().Debug("VoteForPost",
		zap.Int64("userID", userID),
		zap.String("postID", p.PostID),
		zap.Int8("direction", p.Direction))
	
	// 调用Redis处理投票
	return redis.VoteForPost(strconv.Itoa(int(userID)), p.PostID, float64(p.Direction))
}
