package message

import (
	"encoding/json"
	"log/slog"
	"order-service/internal/configuration"
	"order-service/internal/model"
	"order-service/internal/service"
	"order-service/internal/util"

	amqp "github.com/rabbitmq/amqp091-go"
)

type MessageHandler interface {
	RegisterMessagesListener() error
}

type RabbitMQService struct {
	logger       *slog.Logger
	client       MessageQueueClient
	orders       service.OrderService
	validator    util.Validator
	exchangeName string
	exchangeType string
	queue        string
}

func NewMessageHandler(logger *slog.Logger, client MessageQueueClient, orders service.OrderService, validator util.Validator, config configuration.Config) MessageHandler {
	return &RabbitMQService{
		logger:       logger,
		client:       client,
		orders:       orders,
		validator:    validator,
		exchangeName: config.Messaging.Order.Exchange.Name,
		exchangeType: config.Messaging.Order.Exchange.Type,
		queue:        config.Messaging.Order.Queue,
	}
}

func (s *RabbitMQService) RegisterMessagesListener() error {
	deliveries, err := s.client.ReadChannel(s.exchangeName, s.exchangeType, s.queue)
	if err != nil {
		return err
	}

	go s.handleDeliveries(deliveries)

	return nil
}

func (s *RabbitMQService) handleDeliveries(deliveries <-chan amqp.Delivery) {
	for delivery := range deliveries {
		err := s.processOrder(delivery)
		if err != nil {
			s.logger.Error("Failed to process order", "error", err.Error())
		}
	}
	s.client.Done()
	s.client.Shutdown()
}

func (s *RabbitMQService) processOrder(delivery amqp.Delivery) error {
	order := model.Order{}
	json.Unmarshal([]byte(delivery.Body), &order)
	err := s.validator.Validate(order)
	if err != nil {
		return err
	}
	_, err = s.orders.Create(order)
	if err == nil {
		err = delivery.Ack(true)
	}

	return err
}
