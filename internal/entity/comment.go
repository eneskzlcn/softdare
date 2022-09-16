package entity

import (
	"time"
)

type Comment struct {
	ID        string
	PostID    string
	UserID    string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	Username  string
}
type FormattedComment struct {
	ID        string
	PostID    string
	Content   string
	UserID    string
	Username  string
	CreatedAt string
	UpdatedAt string
}
