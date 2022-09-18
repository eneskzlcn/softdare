package postutil

import (
	"github.com/eneskzlcn/softdare/internal/entity"
	"time"
)

func FormatPosts(posts []*entity.Post, timeAgoFormatter func(time.Time) string) []entity.FormattedPost {
	formattedPosts := make([]entity.FormattedPost, 0)
	for _, postPtr := range posts {
		formattedPosts = append(formattedPosts, FormatPost(postPtr, timeAgoFormatter))
	}
	return formattedPosts
}

func FormatPost(post *entity.Post, timeAgoFormatter func(time.Time) string) (formattedPost entity.FormattedPost) {
	formattedPost.CreatedAt = timeAgoFormatter(post.CreatedAt)
	formattedPost.UpdatedAt = timeAgoFormatter(post.UpdatedAt)
	formattedPost.ID = post.ID
	formattedPost.Content = post.Content
	formattedPost.CommentCount = post.CommentCount
	formattedPost.LikeCount = post.LikeCount
	formattedPost.UserID = post.UserID
	formattedPost.Username = post.Username
	return
}
