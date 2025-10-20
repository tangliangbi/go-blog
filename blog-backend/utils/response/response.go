package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// StandardResponse represents a standard API response structure
type StandardResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
}

// Success creates a successful response
func Success(c *gin.Context, message string, data interface{}, meta interface{}) {
	c.JSON(http.StatusOK, StandardResponse{
		Status:  http.StatusOK,
		Message: message,
		Data:    data,
		Meta:    meta,
	})
}

// Error creates an error response
func Error(c *gin.Context, statusCode int, message string, err error) {
	response := StandardResponse{
		Status:  statusCode,
		Message: message,
	}

	if err != nil {
		response.Error = err.Error()
	}

	c.JSON(statusCode, response)
}

// BadRequest creates a 400 bad request response
func BadRequest(c *gin.Context, message string, err error) {
	Error(c, http.StatusBadRequest, message, err)
}

// InternalError creates a 500 internal server error response
func InternalError(c *gin.Context, message string, err error) {
	Error(c, http.StatusInternalServerError, message, err)
}

// NotFound creates a 404 not found response
func NotFound(c *gin.Context, message string) {
	Error(c, http.StatusNotFound, message, nil)
}

// Unauthorized creates a 401 unauthorized response
func Unauthorized(c *gin.Context, message string) {
	Error(c, http.StatusUnauthorized, message, nil)
}
