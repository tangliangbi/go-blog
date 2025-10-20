package assert

import (
	"com.tang.blog/utils/logger"
	"com.tang.blog/utils/response"
	"github.com/gin-gonic/gin"
)

func Assert(c *gin.Context, err error, message string) bool {
	if err != nil {
		logger.Log.Error(message+": ", err)
		response.InternalError(c, message, err)
		return true
	}
	return false
}

func AssertWithCode(c *gin.Context, err error, message string, code int) bool {
	if err != nil {
		logger.Log.Error(message+": ", err)
		response.Error(c, code, message, err)
		return true
	}
	return false
}
