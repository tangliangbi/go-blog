package comments

import (
	"errors"

	"com.tang.blog/pkg/utils/ctl"
	"com.tang.blog/repository/db/dao"
	"com.tang.blog/repository/model"
	"com.tang.blog/types/v1"
	"com.tang.blog/utils/assert"
	"com.tang.blog/utils/response"
	"github.com/gin-gonic/gin"
)

type CommentService struct {
}

func NewCommentService() *CommentService {
	return &CommentService{}
}

func (s *CommentService) InsertComment(ctx *gin.Context) {
	var req types.CommentInsertRequest
	if assert.AssertWithCode(ctx, ctx.ShouldBind(&req), "参数错误", 400) {
		return
	}

	commentDao := dao.NewCommentDao()
	UserInfo, err := ctl.GetUserInfo(ctx.Request.Context())
	if assert.Assert(ctx, err, "获取用户信息失败") {
		return
	}

	err = commentDao.InsertComment(&model.Comment{Content: req.Content, PostID: req.PostId, UserID: UserInfo.Id})
	if assert.Assert(ctx, err, "添加评论失败") {
		return
	}
}

func (s *CommentService) DeleteById(ctx *gin.Context) {
	var req types.CommentDeleteRequest
	if assert.AssertWithCode(ctx, ctx.ShouldBindUri(&req), "参数错误", 400) {
		return
	}

	commentDao := dao.NewCommentDao()
	UserInfo, err := ctl.GetUserInfo(ctx.Request.Context())
	if assert.Assert(ctx, err, "获取用户信息失败") {
		return
	}

	err = commentDao.DeleteById(&model.Comment{ID: req.Id, UserID: UserInfo.Id})
	if err != nil {
		if errors.Is(err, dao.ErrRecordNotFound) {
			response.NotFound(ctx, "评论不存在")
			return
		}
		if errors.Is(err, dao.ErrPermissionDenied) {
			response.Error(ctx, 403, "无权限删除此评论", nil)
			return
		}
		assert.Assert(ctx, err, "删除评论失败")
		return
	}

	response.Success(ctx, "删除评论成功", nil, nil)
}

func (s *CommentService) UpdateById(ctx *gin.Context) {
	var req types.CommentUpdateRequest
	if assert.AssertWithCode(ctx, ctx.ShouldBindUri(&req), "URI参数错误", 400) {
		return
	}

	var commentReq types.CommentUpdateBody
	if assert.AssertWithCode(ctx, ctx.ShouldBindJSON(&commentReq), "JSON参数错误", 400) {
		return
	}

	commentDao := dao.NewCommentDao()
	UserInfo, err := ctl.GetUserInfo(ctx.Request.Context())
	if assert.Assert(ctx, err, "获取用户信息失败") {
		return
	}

	comment := model.Comment{
		ID:      req.Id,
		Content: commentReq.Content,
		UserID:  UserInfo.Id,
	}

	err = commentDao.UpdateById(&comment)
	if assert.Assert(ctx, err, "更新评论失败") {
		return
	}

	response.Success(ctx, "更新评论成功", comment, nil)
}

func (s *CommentService) QueryCommentByPostId(ctx *gin.Context) {
	var req types.CommentQueryRequest
	if assert.AssertWithCode(ctx, ctx.ShouldBind(&req), "参数错误", 400) {
		return
	}
	commentDao := dao.NewCommentDao()
	comments, err := commentDao.QueryByPostId(req.PostId)
	if assert.Assert(ctx, err, "获取评论失败") {
		return
	}

	response.Success(ctx, "获取评论成功", comments, nil)
}
