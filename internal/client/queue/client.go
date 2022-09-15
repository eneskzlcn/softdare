package queue

import (
	"context"
	"github.com/eneskzlcn/softdare/internal/core/logger"
	"time"
)

type RabbitMQClient interface {
	Consume(onReceived chan []byte, consumer string, queue string)
	PushMessage(message any, queue string) error
}
type UserService interface {
	IncreaseUserPostCount(ctx context.Context, userID string, increaseAmount int) (time.Time, error)
	IncreaseUserFollowerOrFollowedCount(ctx context.Context, userID string, increaseAmount int, isFollower bool) (time.Time, error)
}
type PostService interface {
	IncreasePostCommentCount(ctx context.Context, postID string, increaseAmount int) (time.Time, error)
}

type Service interface {
	UserService
	PostService
}
type Client struct {
	client  RabbitMQClient
	logger  logger.Logger
	service Service
}

func New(client RabbitMQClient, logger logger.Logger, service Service) *Client {
	return &Client{client: client, logger: logger, service: service}
}

func (c *Client) Consume(onReceived chan []byte, consumer string, queue string) {
	c.client.Consume(onReceived, consumer, queue)
}

func (c *Client) PushMessage(message any, queue string) error {
	return c.PushMessage(message, queue)
}
