package routes

import (
	"myApp/controller/post"
	"myApp/controller/vote"
	"myApp/middlewares"

	"github.com/gin-gonic/gin"
)

// initPost initializes tPost routes.
func initPost(Router *gin.RouterGroup) {
	PostRouter := Router.Group("post")
	// 列表
	{
		PostRouter.POST("/create", middlewares.JWTAuth(), post.Create)
		PostRouter.DELETE("/delete/:id", middlewares.JWTAuth(), post.Delete)
		PostRouter.PUT("/update/:id", middlewares.JWTAuth(), post.Update)
		PostRouter.GET("/list", post.List)
		PostRouter.GET("/detail", post.Detail)
		PostRouter.POST("/vote", middlewares.JWTAuth(), vote.Vote)
	}
}
