package mysql

import (
	"bluebell/models"
	"database/sql"

	"go.uber.org/zap"
)

// GetCommunityList 获取社区列表
// 返回值:
//   - communityList: 社区列表
//   - err: 可能的错误
func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := "select community_id, community_name from community"
	if err := db.Select(&communityList, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			// 如果数据库中没有社区记录，记录警告日志
			zap.L().Warn("数据库中没有社区记录")
			err = nil
		}
	}
	return
}

// GetCommunityDetailByID 根据ID查询社区详情
// 参数:
//   - id: 社区ID
// 返回值:
//   - community: 社区详情
//   - err: 可能的错误
func GetCommunityDetailByID(id int64) (community *models.CommunityDetail, err error) {
	community = new(models.CommunityDetail)
	sqlStr := `select 
			community_id, community_name, introduction, create_time
			from community 
			where community_id = ?
	`
	if err := db.Get(community, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			// 如果查询不到对应的社区，返回无效ID错误
			err = ErrorInvalidID
		}
	}
	return community, err
}
