package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var namingStrategy = schema.NamingStrategy{
	SingularTable: true,
}

func NewMockService(mockDB gorm.ConnPool) Database {
	gormDB, _ := gorm.Open(postgres.New(postgres.Config{
		Conn: mockDB,
	}), &gorm.Config{
		NamingStrategy: namingStrategy,
	})

	return &service{
		db: gormDB,
	}
}