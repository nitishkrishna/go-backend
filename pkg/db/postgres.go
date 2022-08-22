package db

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresConfig struct {
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
	SSLMode  string
}

func ParsePostgresConfig(envFilepath string) (*PostgresConfig, error) {
	err := godotenv.Load(envFilepath)
	if err != nil {
		return nil, err
	}
	return &PostgresConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}, nil
}

func NewPostgresConnection(config *PostgresConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}
	return db, nil
}
