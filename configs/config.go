package configs

import (
	"github.com/joho/godotenv"
	"os"
	"song-library/internal/pkg/logger"
)

type Config struct {
	DB     DBConfig
	Server ServerConfig
}
type DBConfig struct {
	Username string
	DBName   string
	Host     string
	Password string
	SSLMode  string
	Port     string
}
type ServerConfig struct {
	ExternalApiAddr string
	Host            string
	Port            string
}

func GetConfig() Config {
	if err := godotenv.Load(); err != nil {
		logger.Log.Fatal("Error loading env file | ", err)
	}

	cfg := Config{}

	cfg.DB.Username = os.Getenv("DB.USER")
	cfg.DB.DBName = os.Getenv("DB.DBNAME")
	cfg.DB.Host = os.Getenv("DB.HOST")
	cfg.DB.Password = os.Getenv("DB.PASSWORD")
	cfg.DB.SSLMode = os.Getenv("DB.SSLMODE")
	cfg.DB.Port = os.Getenv("DB.PORT")

	cfg.Server.ExternalApiAddr = os.Getenv("SERVER.EXTERNAL_API")
	cfg.Server.Host = os.Getenv("SERVER.HOST")
	cfg.Server.Port = os.Getenv("SERVER.PORT")

	return cfg
}
