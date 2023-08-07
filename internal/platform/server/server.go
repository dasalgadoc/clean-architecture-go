package server

import (
	"github.com/dasalgadoc/clean-architecture-go/cmd/builder"
	"github.com/dasalgadoc/clean-architecture-go/internal/platform/server/handler/course"
	"github.com/dasalgadoc/clean-architecture-go/internal/platform/server/handler/health"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type EntryPoints struct {
	courseCreator               course.PostCourseCreator
	courseCreatorCommandHandler course.PostCourseCreatorAsync
}

func New(application builder.Application, port string) *http.Server {
	engine := gin.Default()

	entryPoints := EntryPoints{
		courseCreator:               course.NewPostCourseCreator(application.CourseCreator),
		courseCreatorCommandHandler: course.NewPostCourseCreatorAsync(application.CommandBus),
	}

	registerRoutes(entryPoints, engine)

	return &http.Server{
		Addr:         port,
		Handler:      engine,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
}

func registerRoutes(entryPoints EntryPoints, engine *gin.Engine) {
	engine.GET("/health", health.CheckHandler())
	engine.POST("/course", entryPoints.courseCreator.Do)
	engine.POST("/course/async", entryPoints.courseCreatorCommandHandler.Do)
}
