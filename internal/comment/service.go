package comment

import (
	"context"
	"fmt"
	"github.com/eneskzlcn/softdare/internal/core/logger"
	"github.com/eneskzlcn/softdare/internal/entity"
	contextUtil "github.com/eneskzlcn/softdare/internal/util/context"
	"github.com/rs/xid"
	"time"
)

type CommentRepository interface {
	CreateComment(ctx context.Context, input CreateCommentRequest) (time.Time, error)
	GetCommentsByPostID(ctx context.Context, postID string) ([]*entity.Comment, error)
}
type RabbitMQClient interface {
	PushMessage(message any, queue string) error
}
type Service struct {
	logger         logger.Logger
	repo           CommentRepository
	rabbitmqClient RabbitMQClient
}

func NewService(logger logger.Logger, repo CommentRepository, rabbitmqClient RabbitMQClient) *Service {
	if logger == nil {
		fmt.Println("given logger is nil")
		return nil
	}
	if repo == nil {
		logger.Error("given comment repository is nil")
		return nil
	}
	return &Service{logger: logger, repo: repo, rabbitmqClient: rabbitmqClient}
}

func (s *Service) CreateComment(ctx context.Context, in CreateCommentInput) (*entity.Comment, error) {
	in.Prepare()
	if err := in.Validate(); err != nil {
		s.logger.Errorf("validation error. Error: %s", err.Error())
	}
	user, exists := contextUtil.FromContext[entity.User]("user", ctx)
	if !exists {
		s.logger.Error("%s , Exists: %t", entity.CouldNotTakeUserFromContext, exists)
		return nil, entity.CouldNotTakeUserFromContext.Err()
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
	err = s.rabbitmqClient.PushMessage(IncreasePostCommentCountMessage{
		PostID:         in.PostID,
		IncreaseAmount: 1,
	}, "increase-post-comment-count")
	if err != nil {
		s.logger.Error("error publishing increase-post-comment-count message")
	}
	return &entity.Comment{
		ID:        id,
		PostID:    in.PostID,
		UserID:    user.ID,
		Content:   in.Content,
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
		Username:  user.Username,
	}, nil
}
func (s *Service) GetCommentsByPostID(ctx context.Context, postID string) ([]*entity.Comment, error) {
	if _, err := xid.FromString(postID); err != nil {
		s.logger.Error(entity.InvalidPostID)
		return nil, entity.InvalidPostID.Err()
	}
	comments, err := s.repo.GetCommentsByPostID(ctx, postID)
	if err != nil {
		s.logger.Error("error getting comments from repository")
		return nil, err
	}
	return comments, nil
}
