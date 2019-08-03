package db

import (
	"database/sql"
	"ebox-api/internal/config"
	"fmt"
	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

func New (config *config.DBConfig) (*DB, error) {
	dbConnInfo := fmt.Sprintf(`
		host=%s
		port=%d
		user=%s
		password=%s
		dbname=%s
		sslmode=disable`,
		config.Host,
		config.Port,
		config.Username,
		config.Password,
		config.DBName)

	db, err := sql.Open("postgres", dbConnInfo)

	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}
