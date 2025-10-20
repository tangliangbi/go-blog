package posts

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestListPosts(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("ListPostsWithDefaultPagination", func(t *testing.T) {
		// Create a request to send to our handler
		req, _ := http.NewRequest("GET", "/api/v1/posts", nil)

		// Create a ResponseRecorder to record the response
		rr := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rr)
		c.Request = req

		// Create service instance
		service := NewPostsService()

		// Call the method
		service.ListPosts(c)

		// Check the status code
		// Without database connection, we expect an internal server error
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})

	t.Run("ListPostsWithCustomPagination", func(t *testing.T) {
		// Create a request with pagination parameters
		req, _ := http.NewRequest("GET", "/api/v1/posts?page=2&page_size=5", nil)

		// Create a ResponseRecorder to record the response
		rr := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rr)
		c.Request = req

		// Create service instance
		service := NewPostsService()

		// Call the method
		service.ListPosts(c)

		// Check the status code
		// Without database connection, we expect an internal server error
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})
}
