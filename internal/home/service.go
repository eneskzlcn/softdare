package home

import (
	"context"
	"fmt"
	"github.com/eneskzlcn/softdare/internal/post"
	postUtil "github.com/eneskzlcn/softdare/internal/util/post"
	"go.uber.org/zap"
)

type PostService interface {
	GetPosts(ctx context.Context) ([]*post.Post, error)
}
type Service struct {
	logger      *zap.SugaredLogger
	postService PostService
}

func NewService(postService PostService, logger *zap.SugaredLogger) *Service {
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

func (s *Service) GetPosts(ctx context.Context) ([]Post, error) {
	postPtrs, err := s.postService.GetPosts(ctx)
	if err != nil {
		s.logger.Error("oops getting posts from post service", zap.Error(err))
		return nil, err
	}
	posts := make([]Post, 0)
	for _, postPtr := range postPtrs {
		post := Post{
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
