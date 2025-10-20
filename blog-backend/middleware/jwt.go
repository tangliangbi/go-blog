package middleware

import (
	"github.com/gin-gonic/gin"

	"com.tang.blog/pkg/e"
	"com.tang.blog/pkg/utils/ctl"
	"com.tang.blog/pkg/utils/jwt"
	util "com.tang.blog/pkg/utils/jwt"
)

// AuthMiddleware token验证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		code = e.SUCCESS
		accessToken := c.GetHeader("access_token")
		refreshToken := c.GetHeader("refresh_token")
		if accessToken == "" {
			code = e.InvalidParams
			c.JSON(200, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   "Token不能为空",
			})
			c.Abort()
			return
		}
		newAccessToken, newRefreshToken, err := util.ParseRefreshToken(accessToken, refreshToken)
		if err != nil {
			code = e.ErrorAuthCheckTokenFail
		}
		if code != e.SUCCESS {
			c.JSON(200, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   "鉴权失败",
				"error":  err.Error(),
			})
			c.Abort()
			return
		}
		claims, err := util.ParseToken(newAccessToken)
		if err != nil {
			code = e.ErrorAuthCheckTokenFail
			c.JSON(200, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   err.Error(),
			})
			c.Abort()
			return
		}
		SetToken(c, newAccessToken, newRefreshToken)
		c.Request = c.Request.WithContext(ctl.NewContext(c.Request.Context(), &ctl.UserInfo{Id: claims.ID}))
		ctl.InitUserInfo(c.Request.Context())
		c.Next()
	}
}

func SetToken(c *gin.Context, accessToken, refreshToken string) {
	secure := IsHttps(c)
	c.Header(jwt.AccessTokenHeader, accessToken)
	c.Header(jwt.RefreshTokenHeader, refreshToken)
	c.SetCookie(jwt.AccessTokenHeader, accessToken, jwt.MaxAge, "/", "", secure, true)
	c.SetCookie(jwt.RefreshTokenHeader, refreshToken, jwt.MaxAge, "/", "", secure, true)
}

// 判断是否https
func IsHttps(c *gin.Context) bool {
	if c.GetHeader(jwt.HeaderForwardedProto) == "https" || c.Request.TLS != nil {
		return true
	}
	return false
}
