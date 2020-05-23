package config

import (
	"os"
)

type Config struct {
	DatabaseEndpoint string `json:"databaseEndpoint"`
	DatabaseName     string `json:"databaseName"`
	DatabaseUser     string `json:"databaseUser"`
	DatabasePassword string `json:"databasePassword"`
}

func LoadConfig(config *Config) {
	config.DatabaseEndpoint = os.Getenv("DATABASE_ENDPOINT")
	config.DatabaseName = os.Getenv("DATABASE_NAME")
	config.DatabaseUser = os.Getenv("DATABASE_USER")
	config.DatabasePassword = os.Getenv("DATABASE_PASSWORD")
}
