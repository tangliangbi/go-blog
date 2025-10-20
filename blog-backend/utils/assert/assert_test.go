package assert

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAssert(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("NoError", func(t *testing.T) {
		rr := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rr)

		result := Assert(c, nil, "Operation failed")

		assert.False(t, result)
		assert.Equal(t, 200, rr.Code) // No response sent
	})

	t.Run("WithError", func(t *testing.T) {
		rr := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rr)

		err := errors.New("database connection failed")
		result := Assert(c, err, "Operation failed")

		assert.True(t, result)
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Contains(t, rr.Body.String(), `"status":500`)
		assert.Contains(t, rr.Body.String(), `"msg":"Operation failed"`)
		assert.Contains(t, rr.Body.String(), `"error":"database connection failed"`)
	})
}

func TestAssertWithCode(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("NoError", func(t *testing.T) {
		rr := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rr)

		result := AssertWithCode(c, nil, "Operation failed", http.StatusBadRequest)

		assert.False(t, result)
		assert.Equal(t, 200, rr.Code) // No response sent
	})

	t.Run("WithError", func(t *testing.T) {
		rr := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rr)

		err := errors.New("validation failed")
		result := AssertWithCode(c, err, "Invalid input", http.StatusBadRequest)

		assert.True(t, result)
		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Contains(t, rr.Body.String(), `"status":400`)
		assert.Contains(t, rr.Body.String(), `"msg":"Invalid input"`)
		assert.Contains(t, rr.Body.String(), `"error":"validation failed"`)
	})
}
