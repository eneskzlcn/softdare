package post

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/eneskzlcn/softdare/internal/core/logger"
	"github.com/eneskzlcn/softdare/internal/entity"
	"time"
)

type DB interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type Repository struct {
	logger logger.Logger
	db     DB
}

func NewRepository(db DB, logger logger.Logger) *Repository {
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
	r.logger.Debugf("A request for creating new post arrived to post repository with id:%s", request.ID)
	query := `
		INSERT INTO posts (id, user_id, content) 
		VALUES ($1, $2, $3) 
		RETURNING created_at`
	row := r.db.QueryRowContext(ctx, query, request.ID, request.UserID, request.Content)
	var createdAt time.Time
	err := row.Scan(&createdAt)
	return createdAt, err
}

func (r *Repository) GetPosts(ctx context.Context) ([]*entity.Post, error) {
	r.logger.Debug("query arrived for all posts")
	query := `
		SELECT posts.id, posts.user_id, posts.content, posts.created_at, posts.updated_at,posts.comment_count, users.username
		FROM posts 
		INNER JOIN users ON posts.user_id = users.id 
		ORDER BY posts.id DESC`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*entity.Post
	for rows.Next() {
		var i entity.Post
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Content,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.CommentCount,
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
func (r *Repository) GetPostByID(ctx context.Context, postID string) (*entity.Post, error) {
	r.logger.Debug("Query arrived for get post by id :%", postID)

	query := `
		SELECT posts.id, posts.user_id, posts.content, posts.created_at, posts.updated_at, posts.comment_count, users.username
		FROM posts 
		INNER JOIN users ON posts.user_id = users.id 
		WHERE posts.id = $1`

	row := r.db.QueryRowContext(ctx, query, postID)
	var i entity.Post
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Content,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CommentCount,
		&i.Username,
	)
	if err != nil {
		r.logger.Error("error on getting the post from database")
	}
	return &i, err
}
func (r *Repository) IncreasePostCommentCount(ctx context.Context, postID string, increaseAmount int) (time.Time, error) {
	query := `
	UPDATE posts
	SET comment_count = comment_count + $1, updated_at = now()
	WHERE id = $2;`
	row := r.db.QueryRowContext(ctx, query, increaseAmount, postID)
	var updatedAt time.Time
	if err := row.Scan(&updatedAt); err != nil {
		r.logger.Error("error scanning query returning")
		return updatedAt, err
	}
	return updatedAt, nil
}
