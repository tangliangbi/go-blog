package v1

import (
	"com.tang.blog/service/posts"
	"github.com/gin-gonic/gin"
)

var postsService = posts.NewPostsService()

func InsertPostsHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 创建文章
		postsService.InsertPost(ctx)
	}
}

func DeleteByPostsHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		postsService.DeletePosts(ctx)
	}
}

func ListPostsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		postsService.ListPosts(c)
	}
}

func GetPostByIdHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		postsService.GetPostById(c)
	}
}

func UpdatePostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		postsService.UpdatePost(c)
	}
}