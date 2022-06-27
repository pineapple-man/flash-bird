package rabbitmq

import (
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	channel *amqp.Channel
}
