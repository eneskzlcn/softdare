package queue

import (
	"context"
	"encoding/json"
	"github.com/eneskzlcn/softdare/internal/message"
)

func (c *Client) ConsumeCommentCreated() {
	onReceivedChan := make(chan []byte, 0)
	go c.consume(onReceivedChan, "comment-created-consumer", "comment-created")
	var forever chan struct{}
	go func() {
		for d := range onReceivedChan {
			c.logger.Debug("comment-created-consumer received a message")
			var message message.CommentCreated
			err := json.Unmarshal(d, &message)
			if err != nil {
				c.logger.Error("unmarshalling error")
				continue
			}
			_, err = c.service.AdjustPostCommentCount(context.Background(), message.PostID, 1)
			if err != nil {
				c.logger.Error("error on increasing post comment count ", err)
				continue
			}
		}
	}()
	c.logger.Sync()
	<-forever
}

func (c *Client) ConsumePostLikeCreated() {
	onReceivedChan := make(chan []byte, 0)
	go c.consume(onReceivedChan, "post-like-created-consumer", "post-like-created")
	var forever chan struct{}
	go func() {
		for d := range onReceivedChan {
			c.logger.Debug("post-like-created-consumer received a message")
			var message message.PostLikeCreated
			err := json.Unmarshal(d, &message)
			if err != nil {
				c.logger.Error("unmarshalling error")
				continue
			}
			c.logger.Debugf("Message PostID:%s, UserID: %s, CreatedAt:%s", message.PostID, message.UserID, message.CreatedAt.String())
			_, err = c.service.AdjustPostLikeCount(context.Background(), message.PostID, 1)
			if err != nil {
				c.logger.Error("error on adjusting post comment count ", err)
				continue
			}
		}
	}()
	c.logger.Sync()
	<-forever
}
