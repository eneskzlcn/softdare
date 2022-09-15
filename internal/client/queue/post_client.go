package queue

import (
	"context"
	"encoding/json"
	"github.com/eneskzlcn/softdare/internal/entity"
)

func (c *Client) ConsumeIncreasePostCommentCount() {
	onReceivedChan := make(chan []byte, 0)
	go c.client.Consume(onReceivedChan, "increase-post-comment-count-consumer", "increase-post-comment-count")
	var forever chan struct{}
	go func() {
		for d := range onReceivedChan {
			c.logger.Debug("IncreasePostCommentCountConsumer receieved a message")
			var message entity.IncreasePostCommentCountMessage
			err := json.Unmarshal(d, &message)
			if err != nil {
				c.logger.Error("unmarshalling error")
				continue
			}
			_, err = c.service.IncreasePostCommentCount(context.Background(), message.PostID, message.IncreaseAmount)
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
