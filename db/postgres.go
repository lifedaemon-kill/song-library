package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"song-library/configs"
)

func NewDB(cfg configs.DBConfig) (*sqlx.DB, error) {
	psqlInfo := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=%s host=%s",
		cfg.Username,
		cfg.DBName,
		cfg.Password,
		cfg.SSLMode,
		cfg.Host)

	db, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	return db, nil
}
