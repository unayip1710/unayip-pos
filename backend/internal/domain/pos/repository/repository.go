package repository

import (
	database "github.com/unayip1710/unayip-pos/internal/domain/pos/repository/database"
	"github.com/unayip1710/unayip-pos/internal/infrastructure/dependencies"
)

type (
	Repo interface {
	}

	repository struct {
		database database.Database
	}
)

func NewRepository(container *dependencies.Container) Repo {
	return &repository{
		database: database.NewDatabase(container),
	}
}

func NewMockRepository(database database.Database) Repo {
	return &repository{
		database: database,
	}
}
