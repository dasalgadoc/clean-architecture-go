package bootstrap

import (
	"context"
	"github.com/dasalgadoc/clean-architecture-go/cmd/builder"
	"github.com/dasalgadoc/clean-architecture-go/internal/platform/server"
)

type Server interface {
	Run(ctx context.Context) error
}

func Run() error {
	application, err := builder.BuildApplication()
	if err != nil {
		return err
	}

	port := application.Config.Port
	shutdownTimeout := application.Config.ShutdownTimeout

	ctx, srv := server.New(context.Background(), *application, port, shutdownTimeout)
	return srv.Run(ctx)
}
