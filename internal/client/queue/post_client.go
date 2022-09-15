package queue

import (
	"context"
	"encoding/json"
	"github.com/eneskzlcn/softdare/internal/entity"
)

func (c *Client) CommentCreatedConsumer() {
	onReceivedChan := make(chan []byte, 0)
	go c.client.Consume(onReceivedChan, "comment-created-consumer", "comment-created")
	var forever chan struct{}
	go func() {
		for d := range onReceivedChan {
			c.logger.Debug("IncreasePostCommentCountConsumer received a message")
			var message entity.CommentCreatedMessage
			err := json.Unmarshal(d, &message)
			if err != nil {
				c.logger.Error("unmarshalling error")
				continue
			}
			_, err = c.service.IncreasePostCommentCount(context.Background(), message.PostID, 1)
			if err != nil {
				c.logger.Error("error on increasing post comment count ", err)
				//maybe we can add a retry mechanism t
				continue
			}
		}
	}()
	c.logger.Sync()
	<-forever
}
