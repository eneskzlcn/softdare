package entity

import (
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
