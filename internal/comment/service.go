package comment

import (
	"context"
	"fmt"
	"github.com/eneskzlcn/softdare/internal/core/logger"
	contextUtil "github.com/eneskzlcn/softdare/internal/util/context"
	"github.com/rs/xid"
	"time"
)

type CommentRepository interface {
	CreateComment(ctx context.Context, input CreateCommentRequest) (time.Time, error)
	GetCommentsByPostID(ctx context.Context, postID string) ([]*Comment, error)
}
type Service struct {
	logger logger.Logger
	repo   CommentRepository
}

func NewService(logger logger.Logger, repo CommentRepository) *Service {
	if logger == nil {
		fmt.Println("given logger is nil")
		return nil
	}
	if repo == nil {
		logger.Error("given comment repository is nil")
		return nil
	}
	return &Service{logger: logger, repo: repo}
}

func (s *Service) CreateComment(ctx context.Context, in CreateCommentInput) (*Comment, error) {
	in.Prepare()
	if err := in.Validate(); err != nil {
		s.logger.Errorf("validation error. Error: %s", err.Error())
	}
	user, exists := contextUtil.FromContext[User]("user", ctx)
	if !exists {
		s.logger.Error("%s , Exists: %t", ErrCouldNotTakeUserFromContext, exists)
		return nil, ErrCouldNotTakeUserFromContext
	}
	id := xid.New().String()
	createdAt, err := s.repo.CreateComment(ctx, CreateCommentRequest{
		ID:      id,
		PostID:  in.PostID,
		UserID:  user.ID,
		Content: in.Content,
	})
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}
	return &Comment{
		ID:        id,
		PostID:    in.PostID,
		UserID:    user.ID,
		Content:   in.Content,
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
		Username:  user.Username,
	}, nil
}
func (s *Service) GetCommentsByPostID(ctx context.Context, postID string) ([]*Comment, error) {
	if _, err := xid.FromString(postID); err != nil {
		s.logger.Error(ErrInvalidPostID)
		return nil, ErrInvalidPostID
	}
	comments, err := s.repo.GetCommentsByPostID(ctx, postID)
	if err != nil {
		s.logger.Error("error getting comments from repository")
		return nil, err
	}
	return comments, nil
}
