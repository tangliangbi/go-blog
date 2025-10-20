package users

import (
	"com.tang.blog/pkg/utils/jwt"
	dao "com.tang.blog/repository/db/dao"
	"com.tang.blog/repository/model"
	"com.tang.blog/types/v1"
	"com.tang.blog/utils/logger"
	"com.tang.blog/utils/response"
	"github.com/gin-gonic/gin"
)

type UserService struct {
}

type UserRegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

func NewUserService() *UserService {
	return &UserService{}
}

func (service *UserService) Login(ctx *gin.Context) (resp *interface{}, err error) {
	var req types.UserLoginRequest
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		response.Error(ctx, 400, "参数错误", err)
		return
	}

	userDao := dao.NewUserDao()
	user, err := userDao.QueryByUsername(req.Username)
	if err != nil {
		logger.Log.Errorf("查询用户失败: %v", err)
		response.Error(ctx, 500, "查询用户失败", err)
		return
	}

	if user == nil {
		response.Error(ctx, 404, "用户不存在", nil)
		return
	}

	if user.Password != req.Password {
		response.Error(ctx, 401, "密码错误", nil)
		return
	}

	accessToken, refreshToken, err := jwt.GenerateToken(user.ID, user.Username)
	if err != nil {
		logger.Log.Errorf("生成token失败: %v", err)
		response.Error(ctx, 500, "生成token失败", err)
		return
	}

	result := gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}

	var res interface{} = result
	resp = &res

	logger.Log.Infof("用户登录成功: %s", user.Username)

	return
}

func (s *UserService) Register(ctx *gin.Context) {
	var req UserRegisterRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		response.Error(ctx, 400, "参数错误", err)
		return
	}

	user := model.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}

	userDao := dao.NewUserDao()

	err = userDao.InsertUser(&user)
	if err != nil {
		logger.Log.Errorf("用户注册失败: %v", err)
		response.Error(ctx, 500, "用户注册失败", err)
		return
	}

	userResponse := struct {
		ID       uint   `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
	}{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	response.Success(ctx, "用户注册成功", userResponse, nil)
}
