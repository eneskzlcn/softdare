package comment

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
	db     DB
	logger logger.Logger
}

func NewRepository(db DB, logger logger.Logger) *Repository {
	if logger == nil {
		fmt.Println("logger is nil")
		return nil
	}
	if db == nil {
		logger.Error("db is nil")
		return nil
	}
	return &Repository{db: db, logger: logger}
}
func (r *Repository) CreateComment(ctx context.Context, input CreateCommentRequest) (time.Time, error) {
	r.logger.Debugf("Create Comment Request Arrived To Repository With ID:%s", input.ID)

	query := `
		INSERT INTO comments (id, user_id, post_id, content)
		VALUES ($1, $2, $3, $4) 
		RETURNING created_at;`

	row := r.db.QueryRowContext(ctx, query, input.ID, input.UserID, input.PostID, input.Content)
	var createdAt time.Time
	err := row.Scan(&createdAt)
	return createdAt, err

}
func (r *Repository) GetCommentsByPostID(ctx context.Context, postID string) ([]*entity.Comment, error) {
	r.logger.Debugf("Get Comments By Post ID Request Arrived To Repository With Post ID:%s", postID)
	query := `
		SELECT comments.id, comments.user_id, comments.post_id, comments.content, comments.created_at, comments.updated_at, users.username 
		FROM comments
		INNER JOIN users ON comments.user_id = users.id 
		WHERE comments.post_id = $1   
		ORDER BY comments.id DESC;`

	rows, err := r.db.QueryContext(ctx, query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	comments := make([]*entity.Comment, 0)
	for rows.Next() {
		var c entity.Comment
		err = rows.Scan(
			&c.ID,
			&c.UserID,
			&c.PostID,
			&c.Content,
			&c.CreatedAt,
			&c.UpdatedAt,
			&c.Username,
		)
		if err != nil {
			return nil, err
		}
		comments = append(comments, &c)
	}
	if err = rows.Close(); err != nil {
		return nil, err
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return comments, nil
}
