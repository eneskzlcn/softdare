package entity

import "time"

type CommentCreatedMessage struct {
	CommentID string    `json:"comment_id"`
	PostID    string    `json:"post_id"`
	CreatedAt time.Time `json:"created_at"`
}
type PostCreatedMessage struct {
	PostID    string    `json:"post_id"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
type UserFollowCreatedMessage struct {
	FollowerID string    `json:"follower_id"`
	FollowedID string    `json:"followed_id"`
	CreatedAt  time.Time `json:"created_at"`
}
type UserFollowDeletedMessage struct {
	FollowerID string    `json:"follower_id"`
	FollowedID string    `json:"followed_id"`
	DeletedAt  time.Time `json:"create_time"`
}
type PostLikeCreatedMessage struct {
	PostID    string    `json:"post_id"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
type CommentLikeCreatedMessage struct {
	CommentID string    `json:"comment_id"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
