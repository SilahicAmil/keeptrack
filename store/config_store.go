package store

import (
	"changeme/internal/services/models"
	"fmt"
)

func (s *SQLiteStore) SaveUser(user *models.CurrentUser) error {
	fmt.Println("user", user)
	query := `
		INSERT INTO users (id, name, user_id)
			VALUES (?, ?, ?)
		ON CONFLICT(id) DO UPDATE SET
    		name = excluded.name,
   			 user_id = excluded.user_id;
	`

	_, err := s.db.Exec(query, user.DisplayName, user.ID)
	return err
}
