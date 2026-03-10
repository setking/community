package routes

import (
	"myApp/controller/comment"
	"myApp/middlewares"

	"github.com/gin-gonic/gin"
)

// initComment initializes tPost routes.
func initComment(Router *gin.RouterGroup) {
	CommentRouter := Router.Group("comment").Use(middlewares.JWTAuth())
	// 列表
	{
		CommentRouter.POST("/create", comment.Create)
		CommentRouter.GET("/list", comment.List)
	}
}
