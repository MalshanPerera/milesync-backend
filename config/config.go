package config

import (
	"jira-for-peasents/utils"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	User     string
	Password string
	Name     string
	Host     string
	Port     string
}

type Config struct {
	Port string
	DB   DBConfig
}

func (e *Config) Validate() error {
	if e.Port == "" {
		return utils.AppError{Message: "Port is required"}
	}

	if e.DB.User == "" {
		return utils.AppError{Message: "DB User is required"}
	}

	if e.DB.Password == "" {
		return utils.AppError{Message: "DB Password is required"}
	}

	if e.DB.Name == "" {
		return utils.AppError{Message: "DB Name is required"}
	}

	if e.DB.Host == "" {
		return utils.AppError{Message: "DB Host is required"}
	}

	if e.DB.Port == "" {
		return utils.AppError{Message: "DB Port is required"}
	}

	return nil
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	dbConfig := DBConfig{
		User:     os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_DATABASE"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
	}

	return &Config{
		Port: os.Getenv("PORT"),
		DB:   dbConfig,
	}
}
