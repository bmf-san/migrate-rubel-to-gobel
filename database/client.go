package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/bmf-san/migrate-rubel-to-gobel/config"
)

type DB struct {
	GobelConn *sql.DB
	RubelConn *sql.DB
}

func NewClient(cfg *config.Config) *DB {
	gobelDB, err := newDB(cfg.Gobel)
	if err != nil {
		panic(err)
	}

	rubelDB, err := newDB(cfg.Rubel)
	if err != nil {
		panic(err)
	}

	return &DB{
		GobelConn: gobelDB,
		RubelConn: rubelDB,
	}
}

func newDB(cfg config.ConnConfig) (*sql.DB, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.Database)
	conn, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	conn.SetConnMaxLifetime(time.Second * 10)

	return conn, nil
}
