package event

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Emitter struct {
	connection *amqp.Connection
}

func (e *Emitter) setup() error {
	log.Printf("Setup1: %v\n", e.connection)
	channel, err := e.connection.Channel()
	if err != nil {
		return err
	}
	log.Printf("Setup2: %v\n", channel)
	defer channel.Close()
	log.Printf("Setup3: %v\n", channel)
	return declareExchange(channel)
}

func (e *Emitter) Push(event string, severity string) error {
	log.Printf("Push1: %v\n", event)
	channel, err := e.connection.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	log.Printf("Push2: %v\n", channel)
	log.Println("Pushing to channel")

	err = channel.Publish(
		"logs_topic",
		severity,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte(event),
		},
	)
	log.Printf("Push3: %v\n", err)
	if err != nil {
		return err
	}
	log.Printf("Push4: %v\n", err)
	return nil
}

func NewEventEmitter(conn *amqp.Connection) (Emitter, error) {
	log.Printf("NewEventEmitter1: %v\n", conn)
	emitter := Emitter{
		connection: conn,
	}
	log.Printf("NewEventEmitter2: %v\n", emitter)
	err := emitter.setup()
	if err != nil {
		return Emitter{}, err
	}
	log.Printf("NewEventEmitter3: %v\n", emitter)
	return emitter, nil
}