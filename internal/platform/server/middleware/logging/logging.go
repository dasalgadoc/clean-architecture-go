package logging

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		end := time.Now()
		latency := end.Sub(start)

		fmt.Printf("[%s] REQ: %s - %v\n", end.Format(time.RFC822Z), c.Request.URL.Path, latency)
	}
}
