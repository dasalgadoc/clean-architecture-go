package builder

import "time"

// default is a tag from envconfig library
type config struct {
	// Server
	Host            string        `default:"localhost"`
	Port            string        `default:"8081"`
	ShutdownTimeout time.Duration `default:"3s"`

	// Database
	DbUser     string        `default:"root"`
	DbPassword string        `default:"root"`
	DbHost     string        `default:"mysql"` // Docker container name or localhost
	DbPort     string        `default:"3306"`
	DbName     string        `default:"test"`
	DbTimeout  time.Duration `default:"5s"`
}
