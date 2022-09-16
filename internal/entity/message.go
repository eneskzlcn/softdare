package entity

type CommentCreatedMessage struct {
	CommentID string `json:"comment_id"`
	PostID    string `json:"post_id"`
}
type PostCreatedMessage struct {
	PostID string `json:"post_id"`
	UserID string `json:"user_id"`
}
type UserFollowCreatedMessage struct {
	FollowerID string `json:"follower_id"`
	FollowedID string `json:"followed_id"`
}
type UserFollowDeletedMessage struct {
	FollowerID string `json:"follower_id"`
	FollowedID string `json:"followed_id"`
}
