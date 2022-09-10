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
const postsQuery = `SELECT posts.id, posts.user_id, posts.content, posts.created_at, posts.updated_at, users.username
FROM posts INNER JOIN users ON posts.user_id = users.id ORDER BY posts.id DESC`

const postById = `SELECT posts.id, posts.user_id, posts.content, posts.created_at, posts.updated_at, users.username
FROM posts INNER JOIN users ON posts.user_id = users.id WHERE posts.id = $1`

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

func (r *Repository) GetPosts(ctx context.Context) ([]*Post, error) {
	r.logger.Debug("query arrived for all posts")
	rows, err := r.db.QueryContext(ctx, postsQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Content,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Username,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
func (r *Repository) GetPostById(ctx context.Context, postID string) (*Post, error) {
	r.logger.Debug("Query arrived for post", zap.String("post_id", postID))
	row := r.db.QueryRowContext(ctx, postById, postID)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Content,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Username,
	)
	return &i, err
}
