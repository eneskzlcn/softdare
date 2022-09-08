package login

import (
	"context"
	"database/sql"
	"fmt"
	"go.uber.org/zap"
	"time"
)

type DB interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

const createUserQuery = `INSERT INTO users (id, email, username) VALUES ($1, $2, $3) RETURNING created_at`
const userByEmailQuery = `SELECT id, email, username, created_at, updated_at FROM users WHERE email = $1`
const userExistsByEmailQuery = `SELECT EXISTS ( SELECT 1 FROM users WHERE email = $1) `
const userExistsByUsernameQuery = `SELECT EXISTS ( SELECT 1 FROM users WHERE username ILIKE $1 )`

type Repository struct {
	db     DB
	logger *zap.SugaredLogger
}

func NewRepository(logger *zap.SugaredLogger, db DB) *Repository {
	if logger == nil {
		fmt.Println("given logger is nil")
		return nil
	}
	if db == nil {
		logger.Error(ErrDBNil)
		return nil
	}
	return &Repository{db: db, logger: logger}
}

func (r *Repository) CreateUser(ctx context.Context, request CreateUserRequest) (time.Time, error) {
	r.logger.Debug("CREATING NEW USER WITH ", zap.String("User ID", request.ID))
	row := r.db.QueryRowContext(ctx, createUserQuery, request.ID, request.Email, request.Username)
	var createdAt time.Time
	err := row.Scan(&createdAt)
	return createdAt, err
}

func (r *Repository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	row := r.db.QueryRowContext(ctx, userByEmailQuery, email)
	var i User
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
	row := r.db.QueryRowContext(ctx, userExistsByEmailQuery, email)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

func (r *Repository) IsUserExistsByUsername(ctx context.Context, username string) (bool, error) {
	row := r.db.QueryRowContext(ctx, userExistsByUsernameQuery, username)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}
