package models

import "time"

// Post 定义了帖子的数据结构
// 注意：结构体字段的顺序可能会影响内存对齐，从而影响性能
type Post struct {
	ID          int64     `json:"id" db:"post_id"`                                   // 帖子ID
	AuthorID    int64     `json:"author_id" db:"author_id"`                          // 作者ID
	CommunityID int64     `json:"community_id" db:"community_id" binding:"required"` // 社区ID，必填
	Status      int32     `json:"status" db:"status"`                                // 帖子状态
	Title       string    `json:"title" db:"title" binding:"required"`               // 帖子标题，必填
	Content     string    `json:"content" db:"content" binding:"required"`           // 帖子内容，必填
	CreateTime  time.Time `json:"create_time" db:"create_time"`                      // 创建时间
}

// ApiPostDetail 定义了帖子详情API的返回结构
type ApiPostDetail struct {
	AuthorName       string             `json:"author_name"` // 作者名称
	VoteNum          int64              `json:"vote_num"`
	*Post                               // 嵌入帖子结构体
	*CommunityDetail `json:"community"` // 嵌入社区信息
}
