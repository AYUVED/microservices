package event

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/ayuved/microservices-helper/domain"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Emitter struct {
	connection *amqp.Connection
}

func (e *Emitter) setup(name string, types string) error {
	log.Printf("Setup1: %v\n", e.connection)
	channel, err := e.connection.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()
	return declareExchange(name, types, channel)
}

func (e *Emitter) Push(event string, name string, severity string) error {
	channel, err := e.connection.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	err = channel.Publish(
		name,
		severity,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(event),
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (e *Emitter) LogEvent(ctx context.Context, eventEmitter *domain.EventEmitter) error {

	emitter, err := NewLogEventEmitter(e.connection, "log_topic", "topic")
	if err != nil {
		return err
	}
	j, err := json.MarshalIndent(&eventEmitter, "", "\t")
	if err != nil {
		return err
	}
	err = emitter.Push(string(j), "logs_topic", "log.INFO")
	if err != nil {
		return err
	}
	return nil

}


func NewLogEventEmitter(conn *amqp.Connection, name string, types string) (*Emitter, error) {
	log.Printf("NewEventEmitter1: %v\n", conn)
	emitter := Emitter{
		connection: conn,
	}
	log.Printf("NewEventEmitter2: %v\n", emitter)
	err := emitter.setup(name, types)
	
	log.Printf("NewEventEmitter3: %v\n", err)
	if err != nil {
		// return Emitter{}, err
		return nil, fmt.Errorf("db migration error: %v", err)

	}
	log.Printf("NewEventEmitter4: %v\n", emitter)
	return &Emitter{ connection: conn }, nil
}
