package postgres

import (
	"database/sql"
	"kino-site/internal/models"
)

type UserStorage struct {
	DB *sql.DB
}

func (s *UserStorage) CreateUser(email, password string) error {
	_, err := s.DB.Exec(`
	INSERT INTO users (email, password)
	VALUES ($1, $2)
	`, email, password)

	return err
}

func (s *UserStorage) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := s.DB.QueryRow("SELECT id, email, password FROM users WHERE email = $1", email).
		Scan(&user.ID, &user.Email, &user.Password)

	if err != nil {
		return nil, err
	}

	return &user, nil
}