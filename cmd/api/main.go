package main

import (
	"github.com/dasalgadoc/clean-architecture-go/cmd/api/bootstrap"
	"log"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
