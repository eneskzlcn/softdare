package commentutil

import (
	"github.com/eneskzlcn/softdare/internal/entity"
	"time"
)

func FormatComments(comments []*entity.Comment, timeAgoFormatter func(time.Time) string) []entity.FormattedComment {
	formattedComments := make([]entity.FormattedComment, 0)
	for _, comment := range comments {
		formattedComments = append(formattedComments, FormatComment(comment, timeAgoFormatter))
	}
	return formattedComments
}

func FormatComment(comment *entity.Comment, timeAgoFormatter func(time.Time) string) (formattedComment entity.FormattedComment) {
	formattedComment.ID = comment.ID
	formattedComment.PostID = comment.PostID
	formattedComment.Username = comment.Username
	formattedComment.Content = comment.Content
	formattedComment.LikeCount = comment.LikeCount
	formattedComment.CreatedAt = timeAgoFormatter(comment.CreatedAt)
	formattedComment.UpdatedAt = timeAgoFormatter(comment.UpdatedAt)
	return
}
