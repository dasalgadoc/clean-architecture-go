package bootstrap

import "github.com/dasalgadoc/clean-architecture-go/internal/platform/server"

const (
	port = ":8081"
)

func Run() error {
	srv := server.New(port)
	return srv.ListenAndServe()
}
