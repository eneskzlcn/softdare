package queue

import (
	"context"
	"encoding/json"
	"github.com/eneskzlcn/softdare/internal/core/logger"
	"github.com/eneskzlcn/softdare/internal/entity"
	"time"
)

type PostService interface {
	IncreasePostCommentCount(ctx context.Context, postID string, increaseAmount int) (time.Time, error)
}

func (c *Client) IncreasePostCommentCountConsumer(postService PostService, logger logger.Logger) {
	onReceivedChan := make(chan []byte, 0)
	go c.client.Consume(onReceivedChan, "increase-post-comment-count-consumer", "increase-post-comment-count")
	var forever chan struct{}
	go func() {
		for d := range onReceivedChan {
			logger.Debug("IncreasePostCommentCountConsumer receieved a message")
			var message entity.IncreasePostCommentCountMessage
			err := json.Unmarshal(d, &message)
			if err != nil {
				logger.Error("unmarshalling error")
				continue
			}
			_, err = postService.IncreasePostCommentCount(context.Background(), message.PostID, message.IncreaseAmount)
			if err != nil {
				logger.Error("error on increasing post comment count ", err)
				//maybe we can add a retry mechanism t
				continue
			}
		}
	}()
	logger.Sync()
	<-forever
}
