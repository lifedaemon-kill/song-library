package configs

import (
	"github.com/joho/godotenv"
	"os"
	"song-library/logger"
)

type Config struct {
	Username string
	DBName   string
	Host     string
	Password string
	SSLMode  string
	Port     string

	ExternalApi string
}

func GetConfig() Config {
	if err := godotenv.Load(); err != nil {
		logger.Log.Fatal("Error loading env file | ", err)
	}

	cfg := Config{}

	cfg.Username = os.Getenv("USER")
	cfg.DBName = os.Getenv("DBNAME")
	cfg.Host = os.Getenv("HOST")
	cfg.Password = os.Getenv("PASSWORD")
	cfg.SSLMode = os.Getenv("SSLMODE")
	cfg.Port = os.Getenv("PORT")

	cfg.ExternalApi = os.Getenv("EXTERNAL_API")

	return cfg
}
