package models

// gorm is a Go ORM for SQL databases
type Community struct {
	BaseModel
	CommunityId   int64  `gorm:"uniqueIndex;column:community_id;comment '社区ID';"`
	CommunityName string `gorm:"index:idx_community_name;type:varchar(128);column:community_name;comment '社区名称';"`
	Introduction  string `gorm:"type:varchar(256);comment '社区简介';"`
}
type SeqCommunity struct {
	ID int64 `gorm:"primaryKey;column:id;comment:'自增ID'"`
}

// 帖子
type Post struct {
	BaseModel
	PostID      int64  `gorm:"uniqueIndex;column:post_id;comment '帖子ID';"`
	AuthorID    int64  `gorm:"index;column:author_id;comment '作者的用户id';"`
	CommunityID int64  `gorm:"index;column:community_id;comment '所属社区';"`
	Status      int8   `gorm:"type:tinyint(4);default:1;comment '帖子状态';"`
	Title       string `gorm:"type:varchar(128);comment '标题';"`
	Content     string `gorm:"type:varchar(8192);comment '内容';"`
}

// 评论
type Comment struct {
	BaseModel
	CommentID int64  `gorm:"uniqueIndex;column:comment_id;comment '评论ID';"`
	Content   string `gorm:"type:varchar(512);comment '评论内容';"`
	AuthorID  int64  `gorm:"index;column:author_id;comment '评论者的用户id';"`
	PostID    int64  `gorm:"type:bigint(20);column:post_id;comment '所属帖子id';"`
	ParentID  int64  `gorm:"type:bigint(20);column:parent_id;comment '父评论ID';"`
	Status    int8   `gorm:"type:tinyint(4);default:1;comment '评论状态';"`
}
