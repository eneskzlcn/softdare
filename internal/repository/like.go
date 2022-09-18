package repository

import (
	"context"
	"time"
)

func (r *Repository) CreatePostLike(ctx context.Context, userID, postID string) (time.Time, error) {
	query := `
		INSERT INTO post_likes (post_id, user_id)
		VALUES ($1, $2)
		RETURNING created_at;`
	row := r.db.QueryRowContext(ctx, query, postID, userID)
	var createdAt time.Time
	if err := row.Scan(&createdAt); err != nil {
		r.logger.Error(err)
		return time.Time{}, err
	}
	return createdAt, nil
}
