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

func (r *Repository) CreateCommentLike(ctx context.Context, commentID, userID string) (time.Time, error) {
	query := `
		INSERT INTO comment_likes (comment_id, user_id)
		VALUES ($1, $2)
		RETURNING created_at;`
	row := r.db.QueryRowContext(ctx, query, commentID, userID)
	var createdAt time.Time
	if err := row.Scan(&createdAt); err != nil {
		r.logger.Error(err)
		return time.Time{}, err
	}
	return createdAt, nil
}

func (r *Repository) IsPostLikeExists(ctx context.Context, postID, userID string) (bool, error) {
	query := `
	SELECT EXISTS (
		SELECT 1 FROM post_likes pl
		WHERE pl.post_id = $1 AND pl.user_id = $2)`

	row := r.db.QueryRowContext(ctx, query, postID, userID)
	var isPostLikeExists bool
	if err := row.Scan(&isPostLikeExists); err != nil {
		r.logger.Error(err)
		return false, err
	}
	return isPostLikeExists, nil
}

func (r *Repository) IsCommentLikeExists(ctx context.Context, commentID, userID string) (bool, error) {
	query := `
	SELECT EXISTS (
		SELECT 1 FROM comment_likes cl
		WHERE cl.comment_id = $1 AND cl.user_id = $2)`

	row := r.db.QueryRowContext(ctx, query, commentID, userID)
	var isPostLikeExists bool
	if err := row.Scan(&isPostLikeExists); err != nil {
		r.logger.Error(err)
		return false, err
	}
	return isPostLikeExists, nil
}
