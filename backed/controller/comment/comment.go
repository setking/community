package comment

import (
	"myApp/controller/user"
	"myApp/dao"
	"myApp/forms"
	"myApp/middlewares"
	"myApp/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 创建
func Create(c *gin.Context) {
	comment := forms.CommentForm{}
	if err := c.ShouldBind(&comment); err != nil {
		user.HandleValidateError(c, err)
		return
	}
	_userID, err := middlewares.GetCurrentUserID(c)
	if err != nil {
		zap.L().Error("GetCurrentUserID() failed", zap.Error(err))
		utils.ResponseError(c, utils.CodeNotLogin)
		return
	}
	comment.AuthorID, err = utils.Uint64ToInt64Safe(_userID)
	if err != nil {
		zap.L().Error("utils.Uint64ToInt64Safe() failed", zap.Error(err))
		utils.ResponseError(c, utils.CodeServerBusy)
		return
	}
	err = dao.AddComment(comment)
	if err != nil {
		zap.S().Errorf("[AddComment] 新建 【评论失败】 %s", err.Error())
		utils.ResponseError(c, utils.CodeServerBusy)
		return
	}
	utils.ResponseSuccess(c, nil)
}

// 列表
func List(c *gin.Context) {
	pages := c.DefaultQuery("p", "0")
	pagesInt, _ := strconv.Atoi(pages)
	perNums := c.DefaultQuery("n", "0")
	perNumsInt, _ := strconv.Atoi(perNums)
	id := c.DefaultQuery("id", "0")
	postId, _ := strconv.Atoi(id)
	rsp, total, err := dao.ListComment(pagesInt, perNumsInt, int64(postId))
	if err != nil {
		zap.S().Errorf("[ListComment] 查询 【评论列表失败】 %s", err.Error())
		utils.ResponseError(c, utils.CodeServerBusy)
		return
	}
	utils.ResponseSuccess(c, gin.H{"data": rsp, "total": total})
}
