package v1

import (
	"com.tang.blog/service/comments"
	"github.com/gin-gonic/gin"
)

var commentService = comments.NewCommentService()

func InsertCommentHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		commentService.InsertComment(ctx)
	}
}

func DeleteCommentHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		commentService.DeleteById(ctx)
	}
}

func UpdateCommentHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		commentService.UpdateById(ctx)
	}
}

func QueryCommentByPostIdHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		commentService.QueryCommentByPostId(ctx)
	}
}
