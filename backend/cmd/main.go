package main

import (
	"github.com/unayip1710/unayip-pos/internal/infrastructure/configuration/server"
	"github.com/unayip1710/unayip-pos/internal/infrastructure/dependencies"
)

const basePath = "/pos"

func main() {
	container := dependencies.StartDependencies()

	server.NewServer(basePath, container)
}
