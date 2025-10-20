package v1

import (
	"com.tang.blog/service/users"
	"github.com/gin-gonic/gin"
)

var UserService = users.NewUserService()

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		resp, err := UserService.Login(c)
		if resp != nil && err == nil {
			c.JSON(200, resp)
		}
	}
}

func Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		UserService.Register(c)
	}
}
