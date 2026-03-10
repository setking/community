// Package routes handles routing for the application.
package routes

import (
	"myApp/controller/community"
	"myApp/middlewares"

	"github.com/gin-gonic/gin"
)

// InitCommunity initializes community routes.
func initCommunity(Router *gin.RouterGroup) {
	CommunityRouter := Router.Group("community")
	// 列表
	{
		CommunityRouter.POST("/create", middlewares.JWTAuth(), community.Create)
		CommunityRouter.DELETE("/delete/:id", middlewares.JWTAuth(), community.Delete)
		CommunityRouter.PUT("/update/:id", middlewares.JWTAuth(), community.Update)
		CommunityRouter.GET("/list", community.List)
	}
}
