package store

import (
	"database/sql"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite" // pure Go SQLite driver
)

// Thin SQLite Wrapper

var DB *sql.DB

type SQLiteStore struct {
	db *sql.DB
}

func NewSQLiteStore() (*SQLiteStore, error) {
	configDir, err := os.UserConfigDir()

	if err != nil {
		return nil, err
	}

	appDir := filepath.Join(configDir, "keeptrack")

	if err := os.MkdirAll(appDir, os.ModePerm); err != nil {
		return nil, err
	}

	dbPath := filepath.Join(appDir, "keeptrack.db")

	db, err := sql.Open("sqlite", dbPath)

	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(1)

	DB = db

	store := &SQLiteStore{db}

	if err := store.init(); err != nil {
		return nil, err
	}
	return store, nil
}

func (s *SQLiteStore) init() error {
	queries := []string{
		`
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL
    );
    `,
		`
	CREATE TABLE IF NOT EXISTS config (
			id INTEGER PRIMARY KEY CHECK (id = 1),
			pat TEXT NOT NULL,
			org TEXT NOT NULL,
			project TEXT NOT NULL)`,
	}

	for _, q := range queries {
		if _, err := s.db.Exec(q); err != nil {
			return err
		}
	}

	return nil
}
