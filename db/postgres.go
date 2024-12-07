package db

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"song-library/configs"
)

func NewDB(cfg configs.Config) (*sqlx.DB, error) {
	if cfg.Username == "" || cfg.Password == "" || cfg.DBName == "" || cfg.Host == "" || cfg.SSLMode == "" || cfg.Port == "" {
		return nil, errors.New("env config is empty")
	}
	psqlInfo := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=%s host=%s port=%s",
		cfg.Username,
		cfg.DBName,
		cfg.Password,
		cfg.SSLMode,
		cfg.Host,
		cfg.Port)

	db, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
