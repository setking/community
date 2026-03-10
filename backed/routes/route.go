package routes

import (
	"myApp/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 接口路由配置
func Routers(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	Router := gin.Default()
	//跨域配置
	Router.Use(middlewares.Cors())
	ApiGroup := Router.Group("api/v1")
	InitUserRouter(ApiGroup)
	initCommunity(ApiGroup)
	initPost(ApiGroup)
	initComment(ApiGroup)
	Router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return Router
}
