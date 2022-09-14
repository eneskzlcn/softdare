package queue

type RabbitMQClient interface {
	Consume(onReceived chan []byte, consumer string, queue string)
	PushMessage(message any, queue string) error
}
type Client struct {
	client RabbitMQClient
}

func New(client RabbitMQClient) *Client {
	return &Client{client: client}
}

func (c *Client) Consume(onReceived chan []byte, consumer string, queue string) {
	c.client.Consume(onReceived, consumer, queue)
}

func (c *Client) PushMessage(message any, queue string) error {
	return c.PushMessage(message, queue)
}
