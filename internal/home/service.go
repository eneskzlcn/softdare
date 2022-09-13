package home

import (
	"context"
	"fmt"
	"github.com/eneskzlcn/softdare/internal/core/logger"
	"github.com/eneskzlcn/softdare/internal/entity"
	postUtil "github.com/eneskzlcn/softdare/internal/util/post"
)

type PostService interface {
	GetPosts(ctx context.Context) ([]*entity.Post, error)
}
type Service struct {
	logger      logger.Logger
	postService PostService
}

func NewService(postService PostService, logger logger.Logger) *Service {
	if logger == nil {
		fmt.Println("logger is nil")
		return nil
	}
	if postService == nil {
		logger.Error("post service is nil")
		return nil
	}
	return &Service{postService: postService}
}

func (s *Service) GetPosts(ctx context.Context) ([]entity.FormattedPost, error) {
	postPtrs, err := s.postService.GetPosts(ctx)
	if err != nil {
		s.logger.Error("oops getting posts from post service")
		return nil, err
	}
	posts := make([]entity.FormattedPost, 0)
	for _, postPtr := range postPtrs {
		post := entity.FormattedPost{
			ID:           postPtr.ID,
			Username:     postPtr.Username,
			Content:      postPtr.Content,
			CommentCount: postPtr.CommentCount,
		}
		post.CreatedAt = postUtil.FormatPostTime(postPtr.CreatedAt)
		posts = append(posts, post)
	}
	return posts, nil
}
