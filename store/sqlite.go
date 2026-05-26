package store

import (
	"changeme/config"
	"changeme/internal/logger"
	"changeme/internal/services/models"
	"database/sql"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"

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
			provider TEXT NOT NULL,
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

	// TODO: Eventually store PAT into keyring. This is fine for now.

	// Normalize Org/Project/PAT
	pat := strings.TrimSpace(cfg.PAT)
	org := url.PathEscape(strings.TrimSpace(cfg.Org))
	project := url.PathEscape(strings.TrimSpace(cfg.Project))

	// Insert into config
	query := `
		INSERT OR REPLACE INTO config (id, provider, pat, org, project)
		VALUES (?, ?, ?, ?, ?)
		`
	_, err := s.db.Exec(query, 1, cfg.Provider, pat, org, project)

	return err
}

func (s *SQLiteStore) SaveUserToConfig(user *models.CurrentUser) error {
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
			t.State,
			t.Description,
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

func (s *SQLiteStore) GetAppData() (config.AzureCFG, models.CurrentUser, error) {

	logger.Error("test error", "GetAppData")
	// TODO: Get PAT from keyring
	query := `
	SELECT org, project, pat, display_name, email
	FROM config
	LIMIT 1
	`

	var cfg config.AzureCFG
	var user models.CurrentUser

	err := s.db.QueryRow(query).Scan(
		&cfg.Org,
		&cfg.Project,
		&cfg.PAT,
		&user.DisplayName,
		&user.Email,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return config.AzureCFG{}, models.CurrentUser{}, nil
		}
		return config.AzureCFG{}, models.CurrentUser{}, err
	}

	return cfg, user, nil
}

func (s *SQLiteStore) CheckAppState() (bool, config.AzureCFG, models.CurrentUser, error) {
	cfg, user, err := s.GetAppData()
	if err != nil {
		return false, config.AzureCFG{}, models.CurrentUser{}, err
	}

	if cfg.Org == "" || cfg.Project == "" || cfg.PAT == "" {
		return false, config.AzureCFG{}, models.CurrentUser{}, nil
	}

	if user.DisplayName == "" {
		return false, config.AzureCFG{}, models.CurrentUser{}, nil
	}

	return true, cfg, user, nil
}
