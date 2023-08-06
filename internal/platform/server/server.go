package server

import (
	"github.com/dasalgadoc/clean-architecture-go/internal/platform/server/handler/health"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func New(port string) *http.Server {
	engine := gin.Default()
	registerRoutes(engine)

	return &http.Server{
		Addr:         port,
		Handler:      engine,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
}

func registerRoutes(engine *gin.Engine) {
	engine.GET("/health", health.CheckHandler())
}
