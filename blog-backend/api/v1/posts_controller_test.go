package v1

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"com.tang.blog/repository/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestInsertPostsHandler(t *testing.T) {
	// Switch to test mode so you don't get such noisy logs
	gin.SetMode(gin.TestMode)

	t.Run("ValidPost", func(t *testing.T) {
		// Create a sample post
		post := model.Post{
			Title:   "Test Post",
			Content: "This is a test post content",
			UserID:  1,
		}

		// Convert struct to JSON
		postJSON, _ := json.Marshal(post)

		// Create a request to send to our handler
		req, _ := http.NewRequest("POST", "/api/v1/posts", bytes.NewBuffer(postJSON))
		req.Header.Set("Content-Type", "application/json")

		// Create a ResponseRecorder to record the response
		rr := httptest.NewRecorder()
		router := gin.New()
		router.POST("/api/v1/posts", InsertPostsHandler())

		// Perform the request
		router.ServeHTTP(rr, req)

		// Check the status code
		// Without database connection, we expect an internal server error
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})

	t.Run("InvalidPost", func(t *testing.T) {
		// Create an invalid post (missing required fields)
		invalidPost := map[string]interface{}{
			"title": "Test Post",
			// Missing content and userID
		}

		// Convert struct to JSON
		postJSON, _ := json.Marshal(invalidPost)

		// Create a request to send to our handler
		req, _ := http.NewRequest("POST", "/api/v1/posts", bytes.NewBuffer(postJSON))
		req.Header.Set("Content-Type", "application/json")

		// Create a ResponseRecorder to record the response
		rr := httptest.NewRecorder()
		router := gin.New()
		router.POST("/api/v1/posts", InsertPostsHandler())

		// Perform the request
		router.ServeHTTP(rr, req)

		// Print response body for debugging
		t.Logf("Response body: %s", rr.Body.String())

		// Check the status code
		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})
}
