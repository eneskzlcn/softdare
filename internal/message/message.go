package message

import "time"

type CommentCreated struct {
	CommentID string    `json:"comment_id"`
	PostID    string    `json:"post_id"`
	CreatedAt time.Time `json:"created_at"`
}
type PostCreated struct {
	PostID    string    `json:"post_id"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
type UserFollowCreated struct {
	FollowerID string    `json:"follower_id"`
	FollowedID string    `json:"followed_id"`
	CreatedAt  time.Time `json:"created_at"`
}
type UserFollowDeleted struct {
	FollowerID string    `json:"follower_id"`
	FollowedID string    `json:"followed_id"`
	DeletedAt  time.Time `json:"create_time"`
}
type PostLikeCreated struct {
	PostID    string    `json:"post_id"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
type CommentLikeCreated struct {
	CommentID string    `json:"comment_id"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
type UserCreated struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}
