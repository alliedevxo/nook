package db

import (
	"database/sql"
	"fmt"
	"nook/internal/filesystem"

	_ "github.com/mattn/go-sqlite3"
)

type Store struct {
	db *sql.DB
}

func Open() (*Store, error) {
	path, err := filesystem.GetUserConfigDir()
	if err != nil {
		return nil, err
	}

	dsn := path + "?_foreign_keys=on"

	sqlDB, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}

	sqlDB.SetMaxOpenConns(1)

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("ping database: %w", err)
	}

	store := &Store{db: sqlDB}

	if err := store.InitializeTables(); err != nil {
		return nil, err
	}

	return store, nil
}

func (s *Store) Close() error {
	return s.db.Close()
}






