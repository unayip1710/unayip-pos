package server

import (
	"github.com/gin-gonic/gin"

	"github.com/unayip1710/unayip-pos/internal/infrastructure/dependencies"
	"github.com/unayip1710/unayip-pos/internal/infrastructure/roles/read"
	"github.com/unayip1710/unayip-pos/internal/infrastructure/roles/write"
)

func NewServer(basePath string, container *dependencies.Container) {
	r := gin.Default()

	read.NewRead(container).RegisterRoutes(basePath, r)
	write.NewWrite(container).RegisterRoutes(basePath, r)

	err := r.Run()
	if err != nil {
		panic(err.Error())
	}
}
