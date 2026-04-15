package store

import (
	"changeme/config"
	"changeme/internal/services/models"
	"database/sql"
	"fmt"
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
			id INTEGER PRIMARY KEY,
			pat TEXT NOT NULL,
			org TEXT NOT NULL,
			project TEXT NOT NULL,
			email TEXT,
			display_name TEXT)`,
		`
	CREATE TABLE IF NOT EXISTS tickets (
		id INTEGER PRIMARY KEY,
		title TEXT,
		description TEXT,
		state TEXT,
		tags TEXT,
		assigned_to TEXT,
		is_assigned_to_me BOOLEAN,
		changed_date TEXT,
		last_notified_date TEXT
	)`,
	}

	for _, q := range queries {
		if _, err := s.db.Exec(q); err != nil {
			return err
		}
	}

	return nil
}

func (s *SQLiteStore) StoreConfig(cfg config.AzureCFG) error {
	// Insert into config
	query := `
		INSERT OR REPLACE INTO config (id, pat, org, project)
		VALUES (?, ?, ?, ?)
		`
	_, err := s.db.Exec(query, 1, cfg.PAT, cfg.Org, cfg.Project)
	return err
}

func (s *SQLiteStore) SaveUser(user *models.CurrentUser) error {
	fmt.Println("user", user)
	query := `
		UPDATE config
		SET email = ?, display_name = ?
		WHERE id = 1
	`

	_, err := s.db.Exec(query, user.Email, user.DisplayName)
	return err
}

func (s *SQLiteStore) SaveTickets(tickets []models.Ticket) error {
	query := `
		INSERT OR REPLACE INTO tickets (
			id,
			title,
			state,
			description,
			tags,
			assigned_to,
			is_assigned_to_me,
			changed_date,
			last_notified_date
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	for _, t := range tickets {
		_, err := s.db.Exec(
			query,
			t.ID,
			t.Title,
			t.Description,
			t.State,
			t.Tags,
			t.AssignedTo,
			t.IsAssignedToMe,
			t.ChangedDate,
			t.LastNotifiedDate,
		)

		if err != nil {
			return err
		}
	}

	return nil
}
