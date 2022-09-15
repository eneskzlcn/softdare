package repository

import (
	"context"
	"time"
)

func (r *Repository) CreateUserFollow(ctx context.Context, followerID string, followedID string) (time.Time, error) {
	r.logger.Debugf("Follow user request arrived. Follower: %s, Followed: %s", followerID, followedID)
	query := `
		INSERT INTO user_follows (follower_id, followed_id)
		VALUES ($1, $2)
		RETURNING created_at;`

	row := r.db.QueryRowContext(ctx, query, followerID, followedID)
	var createdAt time.Time
	if err := row.Scan(&createdAt); err != nil {
		r.logger.Error(err)
		return time.Time{}, err
	}
	return createdAt, nil
}
func (r *Repository) IsUserFollowExists(ctx context.Context, followerID string, followedID string) (bool, error) {
	query := `
		SELECT EXISTS(
			SELECT 1 FROM user_follows
			WHERE follower_id = $1 AND followed_id = $2
		);`

	row := r.db.QueryRowContext(ctx, query, followerID, followedID)
	var exists bool
	if err := row.Scan(&exists); err != nil {
		r.logger.Error(err)
		return false, err
	}
	return exists, nil
}
