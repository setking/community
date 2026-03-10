package dao

import (
	"myApp/forms"
	"myApp/global"
	"myApp/initialize"
	"myApp/models"
	"myApp/utils"
	"strconv"
)

func AddComment(comment forms.CommentForm) (err error) {
	PostID, err := strconv.ParseInt(comment.PostID, 10, 64)
	if err != nil {
		return err
	}
	ParentID, err := strconv.ParseInt(comment.PostID, 10, 64)
	if err != nil {
		return err
	}
	rsp := &models.Comment{
		Content:  comment.Content,
		PostID:   PostID,
		AuthorID: comment.AuthorID,
		Status:   comment.Status,
		ParentID: ParentID,
	}
	rsp.CommentID, err = initialize.GetID()
	if err != nil {
		return err
	}
	res := global.DB.Create(rsp)
	if res.Error != nil {
		return res.Error
	}
	return
}

// 获取评论列表
func ListComment(page, pageSize int, postId int64) ([]models.Comment, int64, error) {
	var comments []models.Comment
	var total int64
	result := global.DB.Model(&models.Comment{}).Where("post_id = ?", postId).Count(&total)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	if total == 0 {
		return []models.Comment{}, 0, nil // 返回空数组，不是错误
	}
	query := global.DB.Where("post_id = ?", postId)
	query = query.Scopes(utils.Paginate(page, pageSize))
	data := query.Find(&comments)
	if data.Error != nil {
		return nil, 0, result.Error
	}
	return comments, total, nil
}
