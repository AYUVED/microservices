package event

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/ayuved/microservices-helper/adapters"
	"github.com/ayuved/microservices-helper/domain"
	"github.com/ayuved/microservices/listener/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	conn      *amqp.Connection
	queueName string
}

func NewConsumer(conn *amqp.Connection) (Consumer, error) {
	consumer := Consumer{
		conn: conn,
	}

	err := consumer.setup()
	if err != nil {
		return Consumer{}, err
	}

	return consumer, nil
}

func (consumer *Consumer) setup() error {
	channel, err := consumer.conn.Channel()
	if err != nil {
		return err
	}

	return declareExchange(channel)
}

type Payload struct {
	App       string `json:"app"`
	Name      string `json:"name"`
	Data      interface{} `json:"data"`
	ProcessId string      `json:"process_id"`
	Status    string      `json:"status"`
	Type      string      `json:"type"`
	User      string      `json:"user"`

}

func (consumer *Consumer) Listen(topics []string) error {
	ch, err := consumer.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := declareRandomQueue(ch)
	if err != nil {
		return err
	}

	for _, s := range topics {
		ch.QueueBind(
			q.Name,
			s,
			"logs_topic",
			false,
			nil,
		)

		if err != nil {
			return err
		}
	}

	messages, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		return err
	}

	forever := make(chan bool)
	go func() {
		for d := range messages {
			var payload Payload
			_ = json.Unmarshal(d.Body, &payload)
			log.Printf("Payload: %v\n", payload)
			go handlePayload(payload)
		}
	}()

	fmt.Printf("Waiting for message [Exchange, Queue] [logs_topic, %s]\n", q.Name)
	<-forever

	return nil
}

func handlePayload(payload Payload) {
	log.Printf("HandlePayload: %v\n", payload)
	switch payload.Name {
	case "logViaRabbit", "event":
		log.Printf("HandlePayload: %v\n", payload)
		// log whatever we get
		err := logEvent(payload)
		if err != nil {
			log.Println(err)
		}

	case "auth":
		// authenticate

	// you can have as many cases as you want, as long as you write the logic

	default:
		err := logEvent(payload)
		if err != nil {
			log.Println(err)
		}
	}
}

func logEvent(entry Payload) error {
	log.Printf("LogEvent: %v\n", entry)
	//jsonData, _ := json.MarshalIndent(entry, "", "\t")
	logadapter, err := adapters.NewLogServiceAdapter(config.GetLogServiceUrl())
	if err != nil {
		log.Fatalf("Failed to initialize payment stub. Error: %v", err)
	}
	ctx := context.TODO()

	logservice := domain.Logservice{
		App:  "Logentry.App",
		Name: entry.Name,
		Data: entry.Data,
		ProcessId:  entry.ProcessId,
		Status:    "pending",
		Type:      "event",
		User:      "listener",
	}
	log.Printf("Logservice1: %v\n", logservice)
	err = logadapter.AddLog(ctx, &logservice) // Assign the returned value to a variable
	log.Printf("Logservice2: %v\n", err)
	if err != nil {
		//app.errorJSON(w, err)
		return err
	}
	return nil
}
