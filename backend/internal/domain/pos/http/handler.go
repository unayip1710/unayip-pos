package http

import (
	"github.com/unayip1710/unayip-pos/internal/domain/pos/usecase"

	"github.com/unayip1710/unayip-pos/internal/infrastructure/dependencies"
)

type (
	Handler struct {
		UseCase usecase.UseCase
	}
)

func NewHandler(container *dependencies.Container) *Handler {
	return &Handler{
		UseCase: usecase.NewUse(container),
	}
}
