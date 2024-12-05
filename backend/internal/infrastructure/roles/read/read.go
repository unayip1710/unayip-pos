package read

import (
	"github.com/gin-gonic/gin"

	posHandler "github.com/unayip1710/unayip-pos/internal/domain/pos/http"
	"github.com/unayip1710/unayip-pos/internal/infrastructure/dependencies"
)

type read struct {
	container *dependencies.Container
}

func NewRead(container *dependencies.Container) *read {
	return &read{
		container: container,
	}
}

func (read *read) RegisterRoutes(basePath string, r *gin.Engine) {
	handler := posHandler.NewHandler(read.container)

	r.GET("/ping", handler.Ping)
}
