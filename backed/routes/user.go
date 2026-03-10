package routes

import (
	"myApp/controller/user"

	"github.com/gin-gonic/gin"
)

// InitUserRouter 初始化user路由
func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		// 注册
		UserRouter.POST("/signup", user.SignUpHandler)
		// 登录
		UserRouter.POST("/login", user.LoginHandler)

	}
}
