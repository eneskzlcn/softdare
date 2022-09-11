package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/eneskzlcn/softdare/internal/config"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
	"log"
	"time"
)

type Client struct {
	connection *amqp.Connection
	logger     *zap.SugaredLogger
}

func New(config config.RabbitMQ, logger *zap.SugaredLogger) *Client {
	if logger == nil {
		fmt.Println("given logger is nil")
		return nil
	}
	con, err := amqp.Dial(createConnectionUrl(config))
	if err != nil {
		logger.Error("error occurred when connecting to rabbitmq server", zap.Error(err))
		return nil
	}
	ch, err := con.Channel()
	defer ch.Close()
	for _, queue := range config.Queues {
		_, err = ch.QueueDeclare(queue, false, false, false, false, nil)
		if err != nil {
			logger.Error("error when declaring new queue")
			return nil
		}
	}

	return &Client{connection: con, logger: logger}
}
func (c *Client) PushMessage(message any, queue string) error {
	c.logger.Debug("PUSHING MESSAGE TO RABBITMQ", zap.String("queue", queue))
	messageBytes, err := json.Marshal(message)
	if err != nil {
		c.logger.Error("error marshalling the message", zap.Any("message", message))
		return err
	}
	ch, err := c.connection.Channel()
	if err != nil {
		c.logger.Error("error reaching channel from connection")
		return err
	}
	defer ch.Close()
	context, cancelFn := context.WithTimeout(context.Background(), time.Second*40)
	defer cancelFn()
	err = ch.PublishWithContext(context,
		"", queue, false, false,
		amqp.Publishing{
			Headers:     nil,
			ContentType: "text/plain",
			Body:        messageBytes,
		})
	if err != nil {
		c.logger.Error("error occurred when publishing the message ")
		return err
	}
	return nil
}

func (c *Client) Consume(messageReceived chan []byte, consumer string, queue string) {
	ch, err := c.connection.Channel()
	defer ch.Close()
	if err != nil {
		return
	}
	msgs, err := ch.Consume(
		queue,
		consumer,
		true,
		false,
		false,
		false,
		nil,
	)
	var forever chan struct{}
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			messageReceived <- d.Body
		}
	}()
	<-forever
}
func createConnectionUrl(config config.RabbitMQ) string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s/", config.Username, config.Password, config.Host, config.Port)
}
