package db

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"song-library/configs"
)

func NewDB(cfg configs.DBConfig) (*sqlx.DB, error) {
	if cfg.Username == "" || cfg.Password == "" || cfg.DBName == "" || cfg.Host == "" || cfg.SSLMode == "" {
		return nil, errors.New("env config is empty")
	}
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
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
