package queue

import (
	"context"
	"encoding/json"
	"github.com/eneskzlcn/softdare/internal/entity"
)

func (c *Client) ConsumeCommentLikeCreated() {
	onReceivedChan := make(chan []byte, 0)
	go c.consume(onReceivedChan, "comment-like-created-consumer", "comment-like-created")
	var forever chan struct{}
	go func() {
		for d := range onReceivedChan {
			var message entity.CommentLikeCreatedMessage
			c.logger.Debug("comment-like-created-consumer received a message", c.logger.AnyModifier("message", message))
			err := json.Unmarshal(d, &message)
			if err != nil {
				c.logger.Error("unmarshall error on comment like created message")
				continue
			}
			_, err = c.service.AdjustCommentLikeCount(context.Background(), message.CommentID, 1)
			if err != nil {
				c.logger.Error(err)
			}
		}
	}()
	c.logger.Sync()
	<-forever
}
