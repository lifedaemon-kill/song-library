package configs

import (
	"github.com/joho/godotenv"
	"os"
	"song-library/logger"
)

type DBConfig struct {
	Username string
	DBName   string
	Host     string
	Password string
	SSLMode  string
}

func GetConfig() DBConfig {
	if err := godotenv.Load(); err != nil {
		logger.Log.Fatal("Error loading env file | ", err)
	}

	cfg := DBConfig{}

	cfg.Username = os.Getenv("USER")
	cfg.DBName = os.Getenv("DBNAME")
	cfg.Host = os.Getenv("HOST")
	cfg.Password = os.Getenv("PASSWORD")
	cfg.SSLMode = os.Getenv("SSLMODE")

	return cfg
}
