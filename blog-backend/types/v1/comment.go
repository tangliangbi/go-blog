package types

type CommentInsertRequest struct {
	Content string `json:"content" binding:"required"`
	PostId  uint   `json:"postId" binding:"required"`
}

type CommentDeleteRequest struct {
	Id uint `uri:"id" binding:"required"`
}

type CommentQueryRequest struct {
	PostId uint `json:"postId" binding:"required"`
}

type CommentUpdateRequest struct {
	Id uint `uri:"id" binding:"required"`
}

type CommentUpdateBody struct {
	Content string `json:"content" binding:"required"`
}
