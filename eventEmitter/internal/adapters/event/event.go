package event

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func declareExchange(name string, types string, ch *amqp.Channel) error {
	return ch.ExchangeDeclare(
		name, // name
		types,       // type
		true,         // durable?
		false,        // auto-deleted?
		false,        // internal?
		false,        // no-wait?
		nil,          // arguements?
	)
}

func declareRandomQueue(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		"",        // name?
		false,     // durable?
		false,     // delete when unused?
		true,      // exclusive?
		false,     // no-wait?
		nil,       // arguments?
	)
}