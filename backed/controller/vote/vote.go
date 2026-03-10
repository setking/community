package vote

import (
	"fmt"
	"myApp/controller/user"
	"myApp/dao"
	"myApp/forms"
	"myApp/middlewares"
	"myApp/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 创建
func Vote(c *gin.Context) {
	voteForm := forms.VoteForm{}
	if err := c.ShouldBindJSON(&voteForm); err != nil {
		user.HandleValidateError(c, err)
		return
	}
	_userID, err := middlewares.GetCurrentUserID(c)
	if err != nil {
		zap.L().Error("GetCurrentUserID() failed", zap.Error(err))
		utils.ResponseError(c, utils.CodeNotLogin)
		return
	}
	userID, err := utils.Uint64ToInt64Safe(_userID)
	if err != nil {
		utils.ResponseError(c, utils.CodeNotLogin)
		return
	}
	if err := dao.PostVote(voteForm.PostID, fmt.Sprint(userID), voteForm.Direction); err != nil {
		utils.ResponseError(c, utils.CodeServerBusy)
		return
	}
	utils.ResponseSuccess(c, nil)
}
