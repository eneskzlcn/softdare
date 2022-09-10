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
	GetPostById(ctx context.Context, postID string) (*Post, error)
	GetPosts(ctx context.Context) ([]*Post, error)
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
	return s.repo.GetPostById(ctx, postID)
}
