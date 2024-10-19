package mysql

import (
	"bluebell/models"
	"github.com/jmoiron/sqlx"
	"strings"
)

// CreatePost 创建新帖子
// 参数:
//   - p: 帖子信息
// 返回值:
//   - err: 可能的错误
func CreatePost(p *models.Post) (err error) {
	sqlStr := `insert into post(
	post_id, title, content, author_id, community_id)
	values (?, ?, ?, ?, ?)
	`
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return
}

// GetPostById 根据帖子ID获取帖子信息
// 参数:
//   - pid: 帖子ID
// 返回值:
//   - post: 帖子信息
//   - err: 可能的错误
func GetPostById(pid int64) (post *models.Post, err error) {
	post = new(models.Post)
	sqlStr := `select
	post_id, title, content, author_id, community_id, create_time
	from post
	where post_id = ?
	`
	err = db.Get(post, sqlStr, pid)
	return
}

// GetPostList 获取帖子列表
// 参数:
//   - page: 页码
//   - size: 每页数量
// 返回值:
//   - posts: 帖子列表
//   - err: 可能的错误
func GetPostList(page, size int64) (posts []*models.Post, err error) {
	sqlStr := `select
	post_id, title, content, author_id, community_id, create_time
	from post
	ORDER BY create_time DESC
	limit ?,?
	`
	posts = make([]*models.Post, 0, 2)
	err = db.Select(&posts, sqlStr, (page-1)*size, size)
	return
}

// GetPostListByIDs 根据帖子ID列表获取帖子信息
// 参数:
//   - ids: 帖子ID列表
// 返回值:
//   - postList: 帖子列表
//   - err: 可能的错误
func GetPostListByIDs(ids []string) (postList []*models.Post, err error) {
	sqlStr := `select post_id, title, content, author_id, community_id, create_time
	from post
	where post_id in (?)
	order by FIND_IN_SET(post_id, ?)
	`
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}
	query = db.Rebind(query)
	err = db.Select(&postList, query, args...)
	return
}
