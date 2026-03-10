package forms

// 帖子
type PostForm struct {
	AuthorID    int64  `form:"author_id" json:"author_id"`
	CommunityID string `form:"community_id" json:"community_id"`
	PostID      string `form:"post_id" json:"post_id"`
	Status      int8   `form:"status" json:"status"`
	Title       string `form:"title" json:"title" binding:"required"`
	Content     string `form:"content" json:"content" binding:"required"`
}

// 投票数据
type VoteForm struct {
	PostID    string  `form:"post_id" json:"post_id" binding:"required"`
	Direction float64 `form:"direction" json:"direction" binding:"required,custom_direction"` // 赞成票（1）反对票（-1）
}
