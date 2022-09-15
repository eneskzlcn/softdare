package repository

import (
	"context"
	"time"
)

func (r *Repository) FollowUser(ctx context.Context, followerID string, followedID string) (time.Time, error) {
	r.logger.Debugf("Follow user request arrived. Follower: %s, Followed: %s", followerID, followedID)
	query := `
		INSERT INTO user_follows (follower_id, followed_id)
		VALUES $1, $2
		RETURNING created_at;`
	row := r.db.QueryRowContext(ctx, query, followerID, followedID)
	var createdAt time.Time
	if err := row.Scan(createdAt); err != nil {
		r.logger.Error(err)
		return time.Time{}, err
	}
	return createdAt, nil
}
