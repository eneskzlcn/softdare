package queue

import (
	"context"
	"encoding/json"
	"github.com/eneskzlcn/softdare/internal/entity"
	"sync"
)

func (c *Client) ConsumePostCreated() {
	onReceivedChan := make(chan []byte, 0)
	go c.consume(onReceivedChan, "post-created-consumer", "post-created")
	var forever chan struct{}
	for d := range onReceivedChan {
		c.logger.Debug("Post Created Consumer received a message")
		var msg entity.PostCreatedMessage
		if err := json.Unmarshal(d, &msg); err != nil {
			c.logger.Error("unmarshalling error", c.logger.ErrorModifier(err))
			continue
		}
		updatedAt, err := c.service.AdjustUserPostCount(context.Background(), msg.UserID, 1)
		if err != nil {
			c.logger.Error("error increasing user post count from service", c.logger.ErrorModifier(err))
			continue
		}
		c.logger.Debug("Post count of user increased",
			c.logger.StringModifier("userID", msg.UserID),
			c.logger.AnyModifier("increaseAmount", 1),
			c.logger.StringModifier("updatedAt", updatedAt.String()))
	}
	<-forever
}
func (c *Client) ConsumeUserFollowCreated() {
	onReceivedChan := make(chan []byte, 0)
	go c.consume(onReceivedChan, "user-follow-created-consumer", "user-follow-created")
	var forever chan struct{}
	for d := range onReceivedChan {
		var msg entity.UserFollowCreatedMessage
		if err := json.Unmarshal(d, &msg); err != nil {
			c.logger.Error("unmarshalling error", c.logger.ErrorModifier(err))
			continue
		}
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			_, err := c.service.AdjustUserFollowerOrFollowedCount(context.Background(),
				msg.FollowerID, 1, true)
			if err != nil {
				c.logger.Error("error increasing user followed count", c.logger.ErrorModifier(err))
			}
			defer wg.Done()
		}()
		go func() {
			_, err := c.service.AdjustUserFollowerOrFollowedCount(context.Background(),
				msg.FollowedID, 1, false)
			if err != nil {
				c.logger.Error("error increasing user follower count", c.logger.ErrorModifier(err))
			}
			defer wg.Done()
		}()
		wg.Wait()
	}
	<-forever
}
