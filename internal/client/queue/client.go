package queue

import (
	"context"
	"github.com/eneskzlcn/softdare/internal/core/logger"
	"github.com/eneskzlcn/softdare/internal/core/mail"
	"time"
)

type RabbitMQClient interface {
	Consume(onReceived chan []byte, consumer string, queue string)
	PushMessage(message any, queue string) error
}

type UserService interface {
	AdjustUserPostCount(ctx context.Context, userID string, adjustment int) (time.Time, error)
	AdjustUserFollowerOrFollowedCount(ctx context.Context, userID string, adjustment int, isFollower bool) (time.Time, error)
}

type PostService interface {
	AdjustPostCommentCount(ctx context.Context, postID string, adjustment int) (time.Time, error)
	AdjustPostLikeCount(ctx context.Context, postID string, adjustment int) (time.Time, error)
}
type CommentService interface {
	AdjustCommentLikeCount(ctx context.Context, commentID string, adjustment int) (time.Time, error)
}
type Service interface {
	UserService
	PostService
	CommentService
}

type Client struct {
	client      RabbitMQClient
	logger      logger.Logger
	service     Service
	mailService mail.Service
}

func New(client RabbitMQClient, logger logger.Logger, service Service, mailService mail.Service) *Client {
	return &Client{client: client, logger: logger, service: service, mailService: mailService}
}

func (c *Client) consume(onReceived chan []byte, consumer string, queue string) {
	c.client.Consume(onReceived, consumer, queue)
}

func (c *Client) PushMessage(message any, queue string) error {
	return c.client.PushMessage(message, queue)
}
