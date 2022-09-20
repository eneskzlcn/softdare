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

type UserRepository interface {
	IsUserExistsByEmail(ctx context.Context, email string) (bool, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	GetUserByUsername(ctx context.Context, username string) (*entity.User, error)
	IsUserExistsByUsername(ctx context.Context, username string) (bool, error)
	IsUserExistsByID(ctx context.Context, userID string) (bool, error)
	CreateUser(ctx context.Context, userID, email, username string) (time.Time, error)
	AdjustUserPostCount(ctx context.Context, userID string, increaseAmount int) (time.Time, error)
	AdjustUserFollowerCount(ctx context.Context, userID string, increaseAmount int) (time.Time, error)
	AdjustUserFollowedCount(ctx context.Context, userID string, increaseAmount int) (time.Time, error)
}

type PostRepository interface {
	GetPosts(ctx context.Context, userID string) ([]*entity.Post, error)
	CreatePost(ctx context.Context, postID, userID, content string) (time.Time, error)
	GetPostByID(ctx context.Context, postID string) (*entity.Post, error)
	AdjustPostCommentCount(ctx context.Context, postID string, increaseAmount int) (time.Time, error)
	GetPostsOfGivenUsers(ctx context.Context, followedUserIDs []string, maxCount int) ([]*entity.Post, error)
	AdjustPostLikeCount(ctx context.Context, postID string, adjustment int) (time.Time, error)
}

type CommentRepository interface {
	CreateComment(ctx context.Context, commentID, userID, postID, content string) (time.Time, error)
	GetCommentsByPostID(ctx context.Context, postID string) ([]*entity.Comment, error)
	AdjustCommentLikeCount(ctx context.Context, commentID string, adjustment int) (time.Time, error)
}

type FollowRepository interface {
	IsUserFollowExists(ctx context.Context, followerID, followedID string) (bool, error)
	CreateUserFollow(ctx context.Context, followerID, followedID string) (time.Time, error)
	DeleteUserFollow(ctx context.Context, followerID, followedID string) (time.Time, error)
	GetFollowedUsersOfFollower(ctx context.Context, userID string) ([]string, error)
}

type LikeRepository interface {
	CreatePostLike(ctx context.Context, userID, postID string) (time.Time, error)
	CreateCommentLike(ctx context.Context, commentID, userID string) (time.Time, error)
}

type Repository interface {
	UserRepository
	PostRepository
	CommentRepository
	FollowRepository
	LikeRepository
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
