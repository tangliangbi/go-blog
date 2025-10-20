package types

type PostDeleteRequest struct {
	PostId uint `uri:"postId" binding:"required"`
}

type PostQueryRequest struct {
	PostId uint `uri:"postId" binding:"required"`
}

type PostUpdateRequest struct {
	PostId uint `uri:"postId" binding:"required"`
}

type PostCreateRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}
