package service

import (
	"context"
	"github.com/eneskzlcn/softdare/internal/entity"
	postUtil "github.com/eneskzlcn/softdare/internal/util/post"
)

func (s *Service) GetFormattedPosts(ctx context.Context, userID string) ([]entity.FormattedPost, error) {
	posts, err := s.GetPosts(ctx, userID)
	if err != nil {
		s.logger.Error("oops getting posts from post service")
		return nil, err
	}
	formattedPosts := make([]entity.FormattedPost, 0)
	for _, postPtr := range posts {
		formattedPost := entity.FormattedPost{
			ID:           postPtr.ID,
			Username:     postPtr.Username,
			Content:      postPtr.Content,
			CommentCount: postPtr.CommentCount,
		}
		formattedPost.CreatedAt = postUtil.FormatPostTime(postPtr.CreatedAt)
		formattedPosts = append(formattedPosts, formattedPost)
	}
	return formattedPosts, nil
}
