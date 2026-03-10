package community

import (
	"net/http"
	"strconv"

	"myApp/controller/user"
	"myApp/dao"
	"myApp/forms"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 创建
func Create(c *gin.Context) {
	community := forms.CommunityForm{}
	if err := c.ShouldBind(&community); err != nil {
		user.HandleValidateError(c, err)
		return
	}
	err := dao.AddCommunity(community)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError, gin.H{
				"msg": err.Error(),
			})
		zap.S().Errorf("[AddCommunity] 新建 【社区失败】 %s", err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "创建成功",
	})
}

// 删除
func Delete(c *gin.Context) {
	id := c.Param("id")
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	err = dao.DeleteCommunity(i)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError, gin.H{
				"msg": err.Error(),
			})
		zap.S().Errorf("[DeleteCommunity] 删除 【社区失败】 %s", err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "删除成功",
	})
}

// 更新
func Update(c *gin.Context) {
	community := forms.CommunityForm{}
	if err := c.ShouldBind(&community); err != nil {
		user.HandleValidateError(c, err)
		return
	}
	id := c.Param("id")
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	err = dao.UpdateCommunity(community, i)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError, gin.H{
				"msg": err.Error(),
			})
		zap.S().Errorf("[UpdateCommunity] 更新 【社区失败】 %s", err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "更新成功",
	})
}

// 列表
func List(c *gin.Context) {
	pages := c.DefaultQuery("p", "0")
	pagesInt, _ := strconv.Atoi(pages)
	perNums := c.DefaultQuery("n", "0")
	perNumsInt, _ := strconv.Atoi(perNums)
	rsp, total, err := dao.ListCommunity(pagesInt, perNumsInt)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError, gin.H{
				"msg": err.Error(),
			})
		zap.S().Errorf("[ListCommunity] 查询 【社区失败】 %s", err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Data":  rsp,
		"Total": total,
	})
}
