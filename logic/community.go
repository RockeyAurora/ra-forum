package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
)

// GetCommunityList 获取社区列表
// 返回值:
//   - []*models.Community: 社区列表
//   - error: 可能发生的错误
func GetCommunityList() ([]*models.Community, error) {
	// 查询数据库，获取所有的社区信息并返回
	return mysql.GetCommunityList()
}

// GetCommunityDetail 获取社区详细信息
// 参数:
//   - id: 社区ID
// 返回值:
//   - *models.CommunityDetail: 社区详细信息
//   - error: 可能发生的错误
func GetCommunityDetail(id int64) (*models.CommunityDetail, error) {
	// 根据社区ID查询数据库，获取社区详细信息并返回
	return mysql.GetCommunityDetailByID(id)
}
