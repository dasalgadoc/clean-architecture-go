package recovery

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RecoverFromPanicsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("panic: %v\n", err)
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()

		c.Next()
	}
}
