package env

import (
	"os"
)

// Default PORT to set if .env PORT var is empty or can't load
const DEFAULT_PORT_IF_EMPTY = "8000"

// Env config struct
type EnvApp struct {
	// Server Envs
	SERVER_PORT string

	// Database Envs
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
	DB_USER     string
	DB_PASSWORD string
}

// Get the env configuration
func GetEnv() EnvApp {
	// Heroku smell
	port := os.Getenv("PORT")
	if port == "" {
		port = DEFAULT_PORT_IF_EMPTY
	}

	return EnvApp{
		SERVER_PORT: port,
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_NAME:     os.Getenv("DB_NAME"),
		DB_USER:     os.Getenv("DB_USERNAME"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
	}
}
