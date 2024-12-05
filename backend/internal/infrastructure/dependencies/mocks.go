// Package dependencies provider dependencies
package dependencies

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/unayip1710/unayip-pos/internal/infrastructure/configuration/database"
)

func StartMockDependencies() (*Container, sqlmock.Sqlmock) {

	mockdb, mock, _ := sqlmock.New()

	db := database.NewMockService(mockdb)
	return &Container{
		database: db,
	}, mock
}
