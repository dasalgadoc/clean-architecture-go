package server

import (
	"context"
	"fmt"
	"github.com/dasalgadoc/clean-architecture-go/cmd/builder"
	"github.com/dasalgadoc/clean-architecture-go/internal/platform/server/handler/course"
	"github.com/dasalgadoc/clean-architecture-go/internal/platform/server/handler/health"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Server struct {
	Server          *http.Server
	shutdownTimeout time.Duration
}

type EntryPoints struct {
	courseCreator               course.PostCourseCreator
	courseCreatorCommandHandler course.PostCourseCreatorAsync
}

func New(
	ctx context.Context,
	application builder.Application,
	port string,
	shutdownTimeout time.Duration) (context.Context, Server) {

	engine := gin.Default()

	entryPoints := EntryPoints{
		courseCreator:               course.NewPostCourseCreator(application.CourseCreator),
		courseCreatorCommandHandler: course.NewPostCourseCreatorAsync(application.CommandBus),
	}

	srv := Server{
		Server: &http.Server{
			Addr:         port,
			Handler:      engine,
			ReadTimeout:  15 * time.Second,
			WriteTimeout: 15 * time.Second,
		},
		shutdownTimeout: shutdownTimeout,
	}

	registerRoutes(entryPoints, engine)

	return serverContext(ctx), srv
}

func registerRoutes(entryPoints EntryPoints, engine *gin.Engine) {
	engine.GET("/health", health.CheckHandler())
	engine.POST("/course", entryPoints.courseCreator.Do)
	engine.POST("/course/async", entryPoints.courseCreatorCommandHandler.Do)
}

func (s *Server) Run(ctx context.Context) error {
	go func() {
		if err := s.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("server shutdown: ", err)
		}
	}()

	<-ctx.Done()

	s.launchRegressiveCount()

	ctxShutdown, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.Server.Shutdown(ctxShutdown)
}

func (s *Server) launchRegressiveCount() {
	log.Println("server shutting down...")
	for remainingTime := s.shutdownTimeout; remainingTime > 0; remainingTime -= time.Second {
		fmt.Printf("server shutdown in %s\n", remainingTime)
		time.Sleep(time.Second)
	}
}

func serverContext(ctx context.Context) context.Context {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		<-c
		cancel()
	}()

	return ctx
}
