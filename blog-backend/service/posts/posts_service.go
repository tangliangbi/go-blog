package posts

import (
	"com.tang.blog/pkg/utils/ctl"
	dao "com.tang.blog/repository/db/dao"
	"com.tang.blog/repository/model"
	"com.tang.blog/types/v1"
	"com.tang.blog/utils/assert"
	"com.tang.blog/utils/pagination"
	"com.tang.blog/utils/response"
	"github.com/gin-gonic/gin"
)

type PostsService struct {
}

func NewPostsService() *PostsService {
	return &PostsService{}
}

/**
 * @Description: 创建文章
 * @param c
 */
func (s *PostsService) InsertPost(c *gin.Context) {
	var req types.PostCreateRequest
	if err := c.ShouldBind(&req); err != nil {
		response.BadRequest(c, "参数错误", err)
		return
	}

	// 从上下文中获取用户信息
	userInfo, err := ctl.GetUserInfo(c.Request.Context())
	if err != nil {
		response.Unauthorized(c, "用户未登录")
		return
	}

	// 持久化
	postsDao := dao.NewPostsDao(c.Request.Context())

	post := &model.Post{
		UserID:  userInfo.Id, // 使用从token中解析出的用户ID
		Title:   req.Title,
		Content: req.Content,
	}

	err = postsDao.InsertPost(post)
	if err != nil {
		response.InternalError(c, "创建文章失败", err)
		return
	}

	response.Success(c, "文章创建成功", post, nil)
}

func (s *PostsService) GetPostById(c *gin.Context) {
	var req types.PostQueryRequest
	if assert.AssertWithCode(c, c.ShouldBindUri(&req), "参数错误", 400) {
		return
	}

	postsDao := dao.NewPostsDao(c.Request.Context())
	post, err := postsDao.GetPostById(req.PostId)
	if assert.Assert(c, err, "获取文章失败") {
		return
	}

	response.Success(c, "获取文章成功", post, nil)
}

func (s *PostsService) UpdatePost(c *gin.Context) {
	var req types.PostUpdateRequest
	if assert.AssertWithCode(c, c.ShouldBindUri(&req), "URI参数错误", 400) {
		return
	}

	var post model.Post
	if assert.AssertWithCode(c, c.ShouldBindJSON(&post), "JSON参数错误", 400) {
		return
	}

	postsDao := dao.NewPostsDao(c.Request.Context())
	UserInfo, err := ctl.GetUserInfo(c.Request.Context())
	if assert.Assert(c, err, "获取用户信息失败") {
		return
	}

	post.ID = req.PostId
	post.UserID = UserInfo.Id

	err = postsDao.UpdatePost(&post)
	if assert.Assert(c, err, "更新文章失败") {
		return
	}

	response.Success(c, "更新文章成功", post, nil)
}

func (s *PostsService) DeletePosts(ctx *gin.Context) {
	var req types.PostDeleteRequest
	if assert.AssertWithCode(ctx, ctx.ShouldBindUri(&req), "参数错误", 400) {
		return
	}

	postDao := dao.NewPostsDao(ctx.Request.Context())
	UserInfo, err := ctl.GetUserInfo(ctx.Request.Context())
	if assert.Assert(ctx, err, "获取用户信息失败") {
		return
	}
	err = postDao.DeletePostById(&model.Post{ID: req.PostId, UserID: UserInfo.Id})
	if assert.Assert(ctx, err, "删除文章失败") {
		return
	}

	response.Success(ctx, "删除文章成功", nil, nil)
}

func (s *PostsService) ListPosts(c *gin.Context) {
	paginationParams := pagination.ParsePagination(c)

	title := c.Query("title")
	content := c.Query("content")

	postsDao := dao.NewPostsDao(c.Request.Context())

	result, err := postsDao.ListPostsWithPagination(paginationParams, title, content)
	if assert.Assert(c, err, "获取文章列表失败") {
		return
	}

	response.Success(c, "获取文章列表成功", result.Data, result.Pagination)
}
