package models

// 定义请求的参数结构体
const (
	// OrderTime 按时间排序
	OrderTime = "time"
	// OrderScore 按分数排序
	OrderScore = "score"
)

// ParamSignUp 注册请求参数
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`                     // 用户名
	Password   string `json:"password" binding:"required"`                     // 密码
	RePassword string `json:"re_password" binding:"required,eqfield=Password"` // 确认密码
	Email      string `json:"email"`
}

// ParamLogin 登录请求参数
type ParamLogin struct {
	Username string `json:"username" binding:"required"` // 用户名
	Password string `json:"password" binding:"required"` // 密码
}

// ParamVoteData 投票数据
type ParamVoteData struct {
	PostID    string `json:"post_id" binding:"required"`              // 帖子ID
	Direction int8   `json:"direction,string" binding:"oneof=1 0 -1"` // 投票方向：赞成票(1)、反对票(-1)、取消投票(0)
}

// ParamPostList 获取帖子列表参数
type ParamPostList struct {
	Page        int64  `json:"page" binding:"required" form:"page"`   // 页码
	Size        int64  `json:"size" binding:"required" form:"size"`   // 每页数量
	Order       string `json:"order" binding:"required" form:"order"` // 排序方式
	CommunityID int64  `json:"community_id"  form:"community_id"`     //可以为空
}

// CommunityPost List
type ParamCommunityPostList struct {
	ParamPostList
}

type ParamEmailData struct {
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
