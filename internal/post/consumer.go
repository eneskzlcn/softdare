package post

import (
	"context"
	"encoding/json"
	"go.uber.org/zap"
)

type RabbitMQClient interface {
	Consume(onReceived chan []byte, consumer string, queue string)
}
type IncreasePostCommentCountMessage struct {
	PostID         string `json:"post_id"`
	IncreaseAmount int    `json:"increase_amount"`
}

func IncreasePostCommentCountConsumer(client RabbitMQClient, postService PostService, logger *zap.SugaredLogger) {
	onRecievedChan := make(chan []byte, 0)
	go client.Consume(onRecievedChan, "increase-post-comment-count-consumer", "increase-post-comment-count")
	var forever chan struct{}
	go func() {
		for d := range onRecievedChan {
			logger.Debug("IncreasePostCommentCountConsumer receieved a message")
			var message IncreasePostCommentCountMessage
			err := json.Unmarshal(d, &message)
			if err != nil {
				logger.Error("unmarshalling error")
				continue
			}
			_, err = postService.IncreasePostCommentCount(context.Background(), message.PostID, message.IncreaseAmount)
			if err != nil {
				logger.Error("error on increasing post comment count", zap.String("error", err.Error()))
				//maybe we can add a retry mechanism t
				continue
			}
		}
	}()
	logger.Sync()
	<-forever
}
