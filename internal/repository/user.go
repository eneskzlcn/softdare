package repository

import (
	"context"
	"github.com/eneskzlcn/softdare/internal/entity"
	"time"
)

func (r *Repository) CreateUser(ctx context.Context, userID, email, username string) (time.Time, error) {
	r.logger.Debugf("CREATING NEW USER WITH RepositoryID: %s", userID)

	query := `
		INSERT INTO users (id, email, username) 
		VALUES ($1, $2, $3) 
		RETURNING created_at`
	row := r.db.QueryRowContext(ctx, query, userID, email, username)
	var createdAt time.Time
	err := row.Scan(&createdAt)
	return createdAt, err
}

func (r *Repository) GetUserByID(ctx context.Context, userID string) (*entity.User, error) {
	query := `
		SELECT id, email, username, created_at, updated_at 
		FROM users 
		WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, userID)
	var i entity.User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Username,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}
func (r *Repository) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	query := `
		SELECT id, email, username, created_at, updated_at 
		FROM users 
		WHERE username = $1`
	row := r.db.QueryRowContext(ctx, query, username)
	var i entity.User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Username,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}
func (r *Repository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	query := `
		SELECT id, email, username, created_at, updated_at 
		FROM users 
		WHERE email = $1`
	row := r.db.QueryRowContext(ctx, query, email)
	var i entity.User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Username,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

func (r *Repository) IsUserExistsByEmail(ctx context.Context, email string) (bool, error) {
	query := `
	SELECT EXISTS ( 
		SELECT 1 
		FROM users 
		WHERE email = $1) `
	row := r.db.QueryRowContext(ctx, query, email)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

func (r *Repository) IsUserExistsByUsername(ctx context.Context, username string) (bool, error) {
	query := `
	SELECT EXISTS (
		SELECT 1 
		FROM users 
		WHERE username ILIKE $1 )`
	row := r.db.QueryRowContext(ctx, query, username)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}
