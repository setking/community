package forms

// 评论
type CommentForm struct {
	AuthorID  int64  `form:"author_id" json:"author_id"`
	PostID    string `form:"post_id" json:"post_id" binding:"required"`
	ParentID  string `form:"parent_id" json:"parent_id"`
	CommentID string `form:"comment_id" json:"comment_id"`
	Status    int8   `form:"status" json:"status"`
	Content   string `form:"content" json:"content" binding:"required"`
}
