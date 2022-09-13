package entity

import (
	postUtil "github.com/eneskzlcn/softdare/internal/util/post"
	"time"
)

type Post struct {
	ID           string
	UserID       string
	Content      string
	CommentCount int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Username     string
}
type FormattedPost struct {
	ID           string
	UserID       string
	Content      string
	CommentCount int
	CreatedAt    string
	UpdatedAt    string
	Username     string
}

func FormatPost(post *Post) (formattedPost FormattedPost) {
	formattedPost.CreatedAt = postUtil.FormatPostTime(post.CreatedAt)
	formattedPost.UpdatedAt = postUtil.FormatPostTime(post.UpdatedAt)
	formattedPost.ID = post.ID
	formattedPost.Content = post.Content
	formattedPost.CommentCount = post.CommentCount
	formattedPost.UserID = post.UserID
	formattedPost.Username = post.Username
	return
}
