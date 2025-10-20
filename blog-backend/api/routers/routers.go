package routers

import (
	api "com.tang.blog/api/v1"
	"com.tang.blog/middleware"
	"github.com/gin-gonic/gin"
)

// 初始化路由
func SetupRoutes(router *gin.Engine) {
	v1 := router.Group("api/v1")
	{
		// ===========【用户】=================
		v1.POST("/register", api.Register())
		v1.POST("/login", api.Login())

		authed := v1.Group("/") // 需要登陆保护
		authed.Use(middleware.AuthMiddleware())
		// ===========【文章】=================
		authed.POST("/posts", api.InsertPostsHandler())
		authed.GET("/posts", api.ListPostsHandler())
		authed.GET("/posts/:postId", api.GetPostByIdHandler())
		authed.PUT("/posts/:postId", api.UpdatePostHandler())
		authed.DELETE("/posts/:postId", api.DeleteByPostsHandler())

		// ===========【评论】=================
		authed.POST("/comments", api.InsertCommentHandler())
		authed.DELETE("/comments/:id", api.DeleteCommentHandler())
		authed.PUT("/comments/:id", api.UpdateCommentHandler())
		authed.GET("/posts/:postId/comments", api.QueryCommentByPostIdHandler())
	}
}
