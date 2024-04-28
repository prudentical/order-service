package message

import (
	"fmt"
	"log/slog"
	"order-service/internal/configuration"

	amqp "github.com/rabbitmq/amqp091-go"
)

type MessageQueueClient interface {
	ReadChannel(exchangeName string, exchangeType string, queue string) (<-chan amqp.Delivery, error)
	Done()
	Shutdown() error
}

type RabbitMQClient struct {
	tag        string
	done       chan error
	logger     *slog.Logger
	connection *amqp.Connection
	channel    *amqp.Channel
}

func NewMessageQueueClient(config configuration.Config, logger *slog.Logger) (MessageQueueClient, error) {
	var url = fmt.Sprintf("%s://%s:%s@%s:%d",
		config.Messaging.Protocol,
		config.Messaging.User,
		config.Messaging.Password,
		config.Messaging.Host,
		config.Messaging.Port,
	)

	logger.Debug("Obtaining AMQP connection")
	conn, err := amqp.Dial(url)
	if err != nil {
		return &RabbitMQClient{}, err
	}

	logger.Debug("Opening AMQP channel")
	ch, err := conn.Channel()
	if err != nil {
		return &RabbitMQClient{}, err
	}

	return &RabbitMQClient{logger: logger, connection: conn, channel: ch}, nil
}

func (c *RabbitMQClient) ReadChannel(exchangeName string, exchangeType string, queue string) (<-chan amqp.Delivery, error) {
	err := c.setup(exchangeName, exchangeType, queue)

	if err != nil {
		return nil, err
	}

	return c.channel.Consume(
		queue,
		c.tag, // consumerTag,
		false, // autoAck
		false, // exclusive
		false, // noLocal
		false, // noWait
		nil,   // arguments
	)
}

func (c *RabbitMQClient) Shutdown() error {

	if err := c.channel.Cancel(c.tag, true); err != nil {
		return err
	}

	if err := c.connection.Close(); err != nil {
		return err
	}

	defer c.logger.Info("RabbitMQ client successfully shutdown")

	// wait for service users to exit
	return <-c.done
}

func (c *RabbitMQClient) Done() {
	c.done <- nil
}

func (c *RabbitMQClient) setup(exchangeName string, exchangeType string, queue string) error {
	err := c.channel.ExchangeDeclare(
		exchangeName, // name of the exchange
		exchangeType, // type
		true,         // durable
		false,        // delete when complete
		false,        // internal
		false,        // noWait
		nil,          // arguments
	)
	if err != nil {
		return err
	}

	_, err = c.channel.QueueDeclare(
		queue, // name of the queue
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // noWait
		nil,   // arguments
	)
	if err != nil {
		return err
	}

	err = c.channel.QueueBind(
		queue,        // name of the queue
		queue,        // bindingKey
		exchangeName, // sourceExchange
		false,        // noWait
		nil,          // arguments
	)
	return err
}
