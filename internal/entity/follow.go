package entity

import "time"

type UserFollow struct {
	FollowerID string
	FollowedID string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
