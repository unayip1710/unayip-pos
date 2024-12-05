package database

import (
	database "github.com/unayip1710/unayip-pos/internal/infrastructure/configuration/database"
	"github.com/unayip1710/unayip-pos/internal/infrastructure/dependencies"
)

type (
	Database interface {
	}

	databaseDomain struct {
		db database.Database
	}
)

func NewDatabase(container *dependencies.Container) Database {
	return &databaseDomain{
		db: container.Database(),
	}
}
