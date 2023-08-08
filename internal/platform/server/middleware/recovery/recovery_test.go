package recovery

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecoveryMiddleware(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)
	router := gin.New() // gin.Default() already apply a recovery and middleware
	router.Use(RecoverFromPanicsMiddleware())
	router.GET("/test-middleware", func(ctx *gin.Context) {
		panic("something unexpected happened")
	})

	httpRecorder := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/test-middleware", nil)
	require.NoError(t, err)

	assert.NotPanics(t, func() {
		router.ServeHTTP(httpRecorder, req)
	})
}
