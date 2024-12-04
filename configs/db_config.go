package configs

import "os"

type DBConfig struct {
	Username string
	DBName   string
	Host     string
	Password string
	SSLMode  string
}

func GetConfig() DBConfig {
	cfg := DBConfig{}

	cfg.Username = os.Getenv("USER")
	cfg.DBName = os.Getenv("DBNAME")
	cfg.Host = os.Getenv("HOST")
	cfg.Password = os.Getenv("PASSWORD")
	cfg.SSLMode = os.Getenv("SSLMODE")

	return cfg
}
