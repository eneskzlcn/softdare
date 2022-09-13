package entity

import (
	postUtil "github.com/eneskzlcn/softdare/internal/util/post"
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

func FormatComments(comments []*Comment) []FormattedComment {
	formattedComments := make([]FormattedComment, 0)
	for _, comment := range comments {
		formattedComments = append(formattedComments, formatComment(comment))
	}
	return formattedComments
}
func formatComment(comment *Comment) (formattedComment FormattedComment) {
	formattedComment.ID = comment.ID
	formattedComment.PostID = comment.PostID
	formattedComment.Username = comment.Username
	formattedComment.Content = comment.Content
	formattedComment.CreatedAt = postUtil.FormatPostTime(comment.CreatedAt)
	formattedComment.UpdatedAt = postUtil.FormatPostTime(comment.UpdatedAt)
	return
}
