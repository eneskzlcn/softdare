package post

import (
	"context"
	"errors"
	"fmt"
	"github.com/eneskzlcn/softdare/internal/core/logger"
	contextUtil "github.com/eneskzlcn/softdare/internal/util/context"
	"github.com/rs/xid"
	"time"
)

type PostRepository interface {
	CreatePost(ctx context.Context, request CreatePostRequest) (time.Time, error)
	GetPostByID(ctx context.Context, postID string) (*Post, error)
	GetPosts(ctx context.Context) ([]*Post, error)
	IncreasePostCommentCount(ctx context.Context, postID string, increaseAmount int) (time.Time, error)
}
type Service struct {
	logger logger.Logger
	repo   PostRepository
}

func NewService(repo PostRepository, logger logger.Logger) *Service {
	if logger == nil {
		fmt.Sprintf("given logger to service is nil")
		return nil
	}
	if repo == nil {
		logger.Error(ErrPostRepositoryNil.Error())
		return nil
	}
	return &Service{repo: repo, logger: logger}
}

func (s *Service) CreatePost(ctx context.Context, in CreatePostInput) (*CreatePostResponse, error) {
	in.Prepare()
	if err := in.Validate(); err != nil {
		s.logger.Error("validation oops on creating post")
		return nil, err
	}
	user, exists := contextUtil.FromContext[User](userContextKey, ctx)
	if !exists {
		s.logger.Error("unauthorized request user not exist on context")
		return nil, ErrUnauthorized
	}

	id := xid.New().String()
	createPostRequest := CreatePostRequest{
		ID:      id,
		UserID:  user.ID,
		Content: in.Content,
	}
	createdAt, err := s.repo.CreatePost(ctx, createPostRequest)
	if err != nil {
		s.logger.Error("oops creating post on repository")
		return nil, err
	}
	return &CreatePostResponse{
		ID:        id,
		CreatedAt: createdAt,
	}, nil
}
func (s *Service) GetPosts(ctx context.Context) ([]*Post, error) {
	return s.repo.GetPosts(ctx)
}
func (s *Service) GetPostByID(ctx context.Context, postID string) (*Post, error) {
	_, err := xid.FromString(postID)
	if err != nil {
		s.logger.Error("post id validation oops")
		return nil, err
	}
	if err != nil {
		s.logger.Error("oops getting post from repository")
		return nil, err
	}
	return s.repo.GetPostByID(ctx, postID)
}
func (s *Service) IncreasePostCommentCount(ctx context.Context, postID string, increaseAmount int) (time.Time, error) {
	_, err := xid.FromString(postID)
	if err != nil {
		s.logger.Error("not valid postID : %", postID)
		return time.Time{}, err
	}
	if increaseAmount <= 0 || increaseAmount >= 10 {
		s.logger.Error("comment increase amount should be between 1-9 including 1 and 9")
		return time.Time{}, errors.New("increase amount not valid")
	}
	return s.repo.IncreasePostCommentCount(ctx, postID, increaseAmount)
}
