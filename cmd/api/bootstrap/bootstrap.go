package bootstrap

import (
	"context"
	"github.com/dasalgadoc/clean-architecture-go/cmd/builder"
	"github.com/dasalgadoc/clean-architecture-go/internal/platform/server"
	"time"
)

const (
	shutdownTimeout = 10 * time.Second
	port            = ":8081"
)

type Server interface {
	Run(ctx context.Context) error
}

func Run() error {
	application, err := builder.BuildApplication()
	if err != nil {
		return err
	}

	ctx, srv := server.New(context.Background(), *application, port, shutdownTimeout)
	return srv.Run(ctx)
}
