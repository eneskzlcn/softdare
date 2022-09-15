package queue

import (
	"context"
	"encoding/json"
	"github.com/eneskzlcn/softdare/internal/entity"
)

func (c *Client) ConsumeIncreaseUserPostCount() {
	onReceivedChan := make(chan []byte, 0)
	go c.Consume(onReceivedChan, "increase-user-post-count-consumer", "increase-user-post-count")
	var forever chan struct{}
	for d := range onReceivedChan {
		c.logger.Debug("IncreaseUserPostCount Consumer received a message")
		var msg entity.IncreaseUserPostCountMessage
		if err := json.Unmarshal(d, &msg); err != nil {
			c.logger.Error("unmarshalling error", c.logger.ErrorModifier(err))
			continue
		}
		updatedAt, err := c.service.IncreaseUserPostCount(context.Background(), msg.UserID, msg.IncreaseAmount)
		if err != nil {
			c.logger.Error("error increasing user post count from service", c.logger.ErrorModifier(err))
			continue
		}
		c.logger.Debug("Post count of user increased",
			c.logger.StringModifier("userID", msg.UserID),
			c.logger.AnyModifier("increaseAmount", msg.IncreaseAmount),
			c.logger.StringModifier("updatedAt", updatedAt.String()))
	}
	<-forever
}
