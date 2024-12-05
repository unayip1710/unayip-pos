package repository_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/unayip1710/unayip-pos/internal/domain/pos/repository"
	database "github.com/unayip1710/unayip-pos/internal/domain/pos/repository/database/mocks"
	"github.com/unayip1710/unayip-pos/internal/infrastructure/dependencies"
)

func Test_NewRepository(t *testing.T) {
	ass := assert.New(t)

	container := &dependencies.Container{}

	repo := repository.NewRepository(container)

	ass.NotNil(repo)
}

func Test_NewMockRepository(t *testing.T) {
	ass := assert.New(t)

	databaseRepository := database.NewMockDatabase()

	repo := repository.NewMockRepository(databaseRepository)

	ass.NotNil(repo)
}
