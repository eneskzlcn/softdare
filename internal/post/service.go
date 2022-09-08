package post

import (
	"context"
	"fmt"
	contextUtil "github.com/eneskzlcn/softdare/internal/util/context"
	"github.com/rs/xid"
	"go.uber.org/zap"
	"time"
)

type PostRepository interface {
	CreatePost(ctx context.Context, request CreatePostRequest) (time.Time, error)
}
type Service struct {
	logger *zap.SugaredLogger
	repo   PostRepository
}

func NewService(repo PostRepository, logger *zap.SugaredLogger) *Service {
	if logger == nil {
		fmt.Sprintf("given logger to service is nil")
		return nil
	}
	if repo == nil {
		logger.Error(ErrPostRepositoryNil)
		return nil
	}
	return &Service{repo: repo, logger: logger}
}

func (s *Service) CreatePost(ctx context.Context, in CreatePostInput) (*CreatePostResponse, error) {
	in.Prepare()
	if err := in.Validate(); err != nil {
		return nil, err
	}
	_, exists := contextUtil.FromContext[User](userContextKey, ctx)
	if !exists {
		return nil, ErrUnauthorized
	}
	return nil, nil
	id := xid.New().String()
	createPostRequest := CreatePostRequest{
		ID:      id,
		UserID:  "",
		Content: in.Content,
	}
	createdAt, err := s.repo.CreatePost(ctx, createPostRequest)
	if err != nil {
		return nil, err
	}
	return &CreatePostResponse{
		ID:        id,
		CreatedAt: createdAt,
	}, nil
}
