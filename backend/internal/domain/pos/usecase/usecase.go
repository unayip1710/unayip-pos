package usecase

import (
	"github.com/unayip1710/unayip-pos/internal/domain/pos/repository"
	"github.com/unayip1710/unayip-pos/internal/infrastructure/dependencies"
)

type (
	UseCase interface {
	}

	useCase struct {
		repo repository.Repo
	}
)

func NewUse(container *dependencies.Container) UseCase {
	return &useCase{
		repo: repository.NewRepository(container),
	}
}

func NewMockUse(repositoryPos repository.Repo) UseCase {
	return &useCase{
		repo: repositoryPos,
	}
}
