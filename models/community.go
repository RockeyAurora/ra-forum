// Package models 定义了社区相关的数据模型
package models

import "time"

// Community 表示一个社区的基本信息
type Community struct {
	ID   int64  `json:"id" db:"community_id"`     // 社区的唯一标识符
	Name string `json:"name" db:"community_name"` // 社区的名称
}

// CommunityDetail 包含了社区的详细信息
type CommunityDetail struct {
	ID           int64     `json:"id" db:"community_id"`                    // 社区的唯一标识符
	Name         string    `json:"name" db:"community_name"`                // 社区的名称
	Introduction string    `json:"introduction,omitempty" db:"introduction"` // 社区的介绍，可以为空
	CreateTime   time.Time `json:"create_time" db:"create_time"`            // 社区的创建时间
}
