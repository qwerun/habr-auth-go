package user_repository

import (
	"context"
	"fmt"
	"github.com/qwerun/habr-auth-go/internal/models"
)

func (r *Repository) Create(user *models.User) (string, error) {
	query := `
		INSERT INTO users (email, password_hash, nickname)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	var id string
	err := r.explorer.DB.QueryRowContext(
		context.Background(),
		query,
		user.Email,
		user.PasswordHash,
		user.Nickname,
	).Scan(&id)
	if err != nil {
		return "", fmt.Errorf("registration insert error: %w", err)
	}
	return id, nil
}
