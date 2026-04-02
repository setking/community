// Package api defines the comment model.
package forms

// 评论
// User represents body of Comment request and response.
type CommentForm struct {
	// Comment's author_id.
	// Required: true
	AuthorID int64 `form:"author_id" json:"author_id"`
	// Comment's post_id.
	// Required: true
	PostID string `form:"post_id" json:"post_id" binding:"required"`
	// Comment's parent_id.
	// Required: true
	ParentID string `form:"parent_id" json:"parent_id"`
	// Comment's comment_id.
	// Required: true
	CommentID string `form:"comment_id" json:"comment_id"`
	// Comment's status.
	// Required: true
	Status int8 `form:"status" json:"status"`
	// Comment's content.
	// Required: true
	Content string `form:"content" json:"content" binding:"required"`
}
