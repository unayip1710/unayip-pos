package configuration

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	Username     string
	Password     string
	DBName       string
	SSLMode      string
	Host         string
	DatabasePort string
	TimeZone     string
}

func GetDatabaseConfig() string {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	dbConfig := DatabaseConfig{
		Username:     os.Getenv("DB_USER"),
		Password:     os.Getenv("DB_PASSWORD"),
		DBName:       os.Getenv("DB_NAME"),
		Host:         os.Getenv("DB_HOST"),
		DatabasePort: os.Getenv("DB_PORT"),
	}

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v",
		dbConfig.Host,
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.DBName,
		dbConfig.DatabasePort,
	)
}
