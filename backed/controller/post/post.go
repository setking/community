package post

import (
	"net/http"
	"strconv"

	"myApp/controller/user"
	"myApp/dao"
	"myApp/forms"
	"myApp/middlewares"
	"myApp/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 创建
func Create(c *gin.Context) {
	post := forms.PostForm{}
	if err := c.ShouldBind(&post); err != nil {
		user.HandleValidateError(c, err)
		return
	}
	_userID, err := middlewares.GetCurrentUserID(c)
	if err != nil {
		zap.L().Error("GetCurrentUserID() failed", zap.Error(err))
		utils.ResponseError(c, utils.CodeNotLogin)
		return
	}
	post.AuthorID, err = utils.Uint64ToInt64Safe(_userID)
	if err != nil {
		zap.L().Error("utils.Uint64ToInt64Safe() failed", zap.Error(err))
		utils.ResponseError(c, utils.CodeServerBusy)
		return
	}
	err = dao.AddPost(post)
	if err != nil {
		zap.S().Errorf("[AddPost] 新建 【帖子失败】 %s", err.Error())
		utils.ResponseError(c, utils.CodeServerBusy)
		return
	}
	utils.ResponseSuccess(c, nil)
}

// 删除
func Delete(c *gin.Context) {
	id := c.Param("id")
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	err = dao.DeletePost(i)
	if err != nil {
		zap.S().Errorf("[DeleteCommunity] 删除 【社区失败】 %s", err.Error())
		utils.ResponseError(c, utils.CodeServerBusy)
		return
	}
	utils.ResponseSuccess(c, nil)
}

// 更新
func Update(c *gin.Context) {
	post := forms.PostForm{}
	if err := c.ShouldBind(&post); err != nil {
		user.HandleValidateError(c, err)
		return
	}
	id := c.Param("id")
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	err = dao.UpdatePost(post, i)
	if err != nil {
		zap.S().Errorf("[UpdateCommunity] 更新 【社区失败】 %s", err.Error())
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
	rsp, total, err := dao.ListPost(pagesInt, perNumsInt)
	if err != nil {
		zap.S().Errorf("[ListCommunity] 查询 【社区失败】 %s", err.Error())
		utils.ResponseError(c, utils.CodeServerBusy)
		return
	}
	utils.ResponseSuccess(c, gin.H{"data": rsp, "total": total})
}

func Detail(c *gin.Context) {
	id := c.DefaultQuery("id", "0")
	postId, _ := strconv.Atoi(id)
	rsp, err := dao.GetPostDetail(int64(postId))
	if err != nil {
		zap.S().Errorf("[GetPostDetail] 查询 【贴子详情失败】 %s", err.Error())
		utils.ResponseError(c, utils.CodeServerBusy)
		return
	}
	userData, err := dao.GetUserByID(rsp.AuthorID)
	if err != nil {
		zap.S().Errorf("[GetUserByID] 查询 【用户信息失败】 %s", err.Error())
		utils.ResponseError(c, utils.CodeServerBusy)
		return
	}
	communityData, err := dao.GetCommunityByID(rsp.CommunityID)
	if err != nil {
		zap.S().Errorf("[GetCommunityByID] 查询 【社区信息失败】 %s", err.Error())
		utils.ResponseError(c, utils.CodeServerBusy)
		return
	}
	utils.ResponseSuccess(c, gin.H{
		"data": gin.H{
			"post_id":      rsp.PostID,
			"title":        rsp.Title,
			"content":      rsp.Content,
			"author_id":    userData.UserID,
			"community_id": rsp.CommunityID,
			"created_at":   rsp.CreatedAt.Format("2006-01-02 15:04:05"),
			"author_name":  userData.UserName,
			"community":    communityData.CommunityName,
		},
	})
}
