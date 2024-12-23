package write

import (
	"github.com/gin-gonic/gin"

	"github.com/unayip1710/unayip-pos/internal/infrastructure/dependencies"
)

type write struct {
	container *dependencies.Container
}

func NewWrite(container *dependencies.Container) *write {
	return &write{
		container: container,
	}
}

func (write *write) RegisterRoutes(basePath string, r *gin.Engine) {
}
