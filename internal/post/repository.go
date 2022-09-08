package post

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

const createPostQuery = `INSERT INTO posts (id, user_id, content) VALUES ($1, $2, $3) RETURNING created_at`

type Repository struct {
	logger *zap.SugaredLogger
	db     DB
}

func NewRepository(db DB, logger *zap.SugaredLogger) *Repository {
	if logger == nil {
		fmt.Sprintf("given logger is nil")
		return nil
	}
	if db == nil {
		logger.Error(ErrDatabaseNil)
		return nil
	}
	return &Repository{db: db, logger: logger}
}
func (r *Repository) CreatePost(ctx context.Context, request CreatePostRequest) (time.Time, error) {
	r.logger.Debug("A request for creating new post arrived to post repository", zap.String("id", request.ID))
	row := r.db.QueryRowContext(ctx, createPostQuery, request.ID, request.UserID, request.Content)
	var createdAt time.Time
	err := row.Scan(&createdAt)
	return createdAt, err
}
