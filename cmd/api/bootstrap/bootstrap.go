package bootstrap

import (
	"github.com/dasalgadoc/clean-architecture-go/cmd/builder"
	"github.com/dasalgadoc/clean-architecture-go/internal/platform/server"
)

const (
	port = ":8081"
)

func Run() error {
	application, err := builder.BuildApplication()
	if err != nil {
		return err
	}
	srv := server.New(*application, port)
	return srv.ListenAndServe()
}
