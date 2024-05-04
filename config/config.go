package config

import (
	err_pkg "jira-for-peasants/errors"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type DBConfig struct {
	User         string
	Password     string
	DatabaseName string
	Host         string
	Port         string
}

type AuthConfig struct {
	Secret string
	Expiry string
}

type Config struct {
	Port string
	DB   DBConfig
	Auth AuthConfig
}

func (e *Config) Validate() error {
	if e.Port == "" {
		return err_pkg.AppError{Message: "Port is required"}
	}

	if e.DB.User == "" {
		return err_pkg.AppError{Message: "DB User is required"}
	}

	if e.DB.Password == "" {
		return err_pkg.AppError{Message: "DB Password is required"}
	}

	if e.DB.DatabaseName == "" {
		return err_pkg.AppError{Message: "DB Name is required"}
	}

	if e.DB.Host == "" {
		return err_pkg.AppError{Message: "DB Host is required"}
	}

	if e.DB.Port == "" {
		return err_pkg.AppError{Message: "DB Port is required"}
	}

	if e.Auth.Secret == "" {
		return err_pkg.AppError{Message: "Auth Secret is required"}
	}

	if e.Auth.Expiry == "" {
		return err_pkg.AppError{Message: "Auth Expiry is required"}
	}

	return nil
}

func NewConfig() *Config {
	dbConfig := DBConfig{
		User:         os.Getenv("DB_USERNAME"),
		Password:     os.Getenv("DB_PASSWORD"),
		DatabaseName: os.Getenv("DB_DATABASE"),
		Host:         os.Getenv("DB_HOST"),
		Port:         os.Getenv("DB_PORT"),
	}

	authConfig := AuthConfig{
		Secret: os.Getenv("JWT_SECRET"),
		Expiry: os.Getenv("JWT_EXPIRED"),
	}

	return &Config{
		Port: os.Getenv("PORT"),
		DB:   dbConfig,
		Auth: authConfig,
	}
}
