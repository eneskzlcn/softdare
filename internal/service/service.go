package service

import (
	"context"
	"fmt"
	"github.com/eneskzlcn/softdare/internal/core/logger"
	"github.com/eneskzlcn/softdare/internal/core/session"
	"github.com/eneskzlcn/softdare/internal/entity"
	"time"
)

type RabbitMQClient interface {
	PushMessage(message any, queue string) error
}

type Repository interface {
	IsUserExistsByEmail(ctx context.Context, email string) (bool, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	GetUserByUsername(ctx context.Context, username string) (*entity.User, error)
	IsUserExistsByUsername(ctx context.Context, username string) (bool, error)
	CreateUser(ctx context.Context, userID, email, username string) (time.Time, error)
	GetPosts(ctx context.Context, userID string) ([]*entity.Post, error)
	CreateComment(ctx context.Context, commentID, userID, postID, content string) (time.Time, error)
	GetCommentsByPostID(ctx context.Context, postID string) ([]*entity.Comment, error)
	CreatePost(ctx context.Context, postID, userID, content string) (time.Time, error)
	GetPostByID(ctx context.Context, postID string) (*entity.Post, error)
	IncreasePostCommentCount(ctx context.Context, postID string, increaseAmount int) (time.Time, error)
}
type Service struct {
	logger         logger.Logger
	repository     Repository
	session        session.Session
	rabbitmqClient RabbitMQClient
}

func New(repository Repository, logger logger.Logger, session session.Session, rabbitmqClient RabbitMQClient) *Service {
	if logger == nil {
		fmt.Println(entity.NilLogger)
		return nil
	}
	if repository == nil || session == nil || rabbitmqClient == nil {
		logger.Error(entity.InvalidConstructorArguments)
		return nil
	}
	return &Service{repository: repository, logger: logger, session: session, rabbitmqClient: rabbitmqClient}
}
