package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/ayuved/microservices/broker/config"
	"github.com/ayuved/microservices/broker/internal/adapters/order"
	"github.com/ayuved/microservices/broker/internal/application/core/domain"
)

// RequestPayload describes the JSON that this service accepts as an HTTP Post request
type RequestPayload struct {
	Action string       `json:"action"`
	Order  OrderPayload `json:"order,omitempty"`
	Auth   AuthPayload  `json:"auth,omitempty"`
	Log    LogPayload   `json:"log,omitempty"`
	Mail   MailPayload  `json:"mail,omitempty"`
}

// OrderPayload is the embedded type (in RequestPayload) that describes an email message to be sent
type OrderPayload struct {
	CustomerID int64              `json:"customerID"`
	Status     string             `json:"status"`
	OrderItems []OrderItemPayload `json:"orderItems"`
}

// OrderItemPayload is the embedded type (in RequestPayload) that describes an email message to be sent
type OrderItemPayload struct {
	ProductCode string  `json:"productCode"`
	UnitPrice   float32 `json:"unitPrice"`
	Quantity    int32   `json:"quantity"`
}

// MailPayload is the embedded type (in RequestPayload) that describes an email message to be sent
type MailPayload struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

// AuthPayload is the embedded type (in RequestPayload) that describes an authentication request
type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LogPayload is the embedded type (in RequestPayload) that describes a request to log something
type LogPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

// Broker is a test handler, just to make sure we can hit the broker from a web client
func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}

// HandleSubmission is the main point of entry into the broker. It accepts a JSON
// payload and performs an action based on the value of "action" in that JSON.
func (app *Config) HandleSubmission(w http.ResponseWriter, r *http.Request) {
	var requestPayload RequestPayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	switch requestPayload.Action {
	// case "auth":
	// 	app.authenticate(w, requestPayload.Auth)
	// case "log":
	// 	app.logItemViaRPC(w, requestPayload.Log)
	// case "mail":
	// 	app.sendMail(w, requestPayload.Mail)
	case "order":
		app.PlaceOrder(w, requestPayload.Order)
	default:
		app.errorJSON(w, errors.New("unknown action"))
	}
}

func (app *Config) PlaceOrder(w http.ResponseWriter, o OrderPayload) {
	log.Printf("Order1: %v\n", o)

	// var orderPayload OrderPayload
	orderdapter, err := order.NewAdapter(config.GetOrderServiceUrl())
	if err != nil {
		log.Fatalf("Failed to initialize payment stub. Error: %v", err)
	}
	// convert orderPayload to a format that the order service can understand
	ctx := context.TODO()
	var orderItems []domain.OrderItem
	for _, orderItem := range o.OrderItems {
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice:   orderItem.UnitPrice,
			Quantity:    orderItem.Quantity,
		})
	}
	order := domain.Order{
		CustomerID: o.CustomerID,
		Status:     o.Status,
		OrderItems: orderItems,
	}
	err = orderdapter.Order(ctx, &order) // Assign the returned value to a variable
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	// err := app.readJSON(w, r, &orderPayload)
	// if err != nil {
	// 	app.errorJSON(w, err)
	// 	return
	// }

	// conn, err := grpc.Dial("logger-service:50001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	// if err != nil {
	// 	app.errorJSON(w, err)
	// 	return
	// }
	// defer conn.Close()

	// c := logs.NewLogServiceClient(conn)
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()

	// _, err = c.WriteLog(ctx, &logs.LogRequest{
	// 	LogEntry: &logs.Log{
	// 		Name: requestPayload.Log.Name,
	// 		Data: requestPayload.Log.Data,
	// 	},
	// })
	// if err != nil {
	// 	app.errorJSON(w, err)
	// 	return
	// }

	var payload jsonResponse
	payload.Error = false
	payload.Message = "logged"

	app.writeJSON(w, http.StatusAccepted, payload)
}

// logItem logs an item by making an HTTP Post request with a JSON payload, to the logger microservice
func (app *Config) logItem(w http.ResponseWriter, entry LogPayload) {
	jsonData, _ := json.MarshalIndent(entry, "", "\t")

	logServiceURL := "http://logger-service/log"

	request, err := http.NewRequest("POST", logServiceURL, bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, err)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "logged"

	app.writeJSON(w, http.StatusAccepted, payload)

}

// // authenticate calls the authentication microservice and sends back the appropriate response
// func (app *Config) authenticate(w http.ResponseWriter, a AuthPayload) {
// 	// create some json we'll send to the auth microservice
// 	jsonData, _ := json.MarshalIndent(a, "", "\t")

// 	// call the service
// 	request, err := http.NewRequest("POST", "http://authentication-service/authenticate", bytes.NewBuffer(jsonData))
// 	if err != nil {
// 		app.errorJSON(w, err)
// 		return
// 	}

// 	client := &http.Client{}
// 	response, err := client.Do(request)
// 	if err != nil {
// 		app.errorJSON(w, err)
// 		return
// 	}
// 	defer response.Body.Close()

// 	// make sure we get back the correct status code
// 	if response.StatusCode == http.StatusUnauthorized {
// 		app.errorJSON(w, errors.New("invalid credentials"))
// 		return
// 	} else if response.StatusCode != http.StatusAccepted {
// 		app.errorJSON(w, errors.New("error calling auth service"))
// 		return
// 	}

// 	// create a variable we'll read response.Body into
// 	var jsonFromService jsonResponse

// 	// decode the json from the auth service
// 	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
// 	if err != nil {
// 		app.errorJSON(w, err)
// 		return
// 	}

// 	if jsonFromService.Error {
// 		app.errorJSON(w, err, http.StatusUnauthorized)
// 		return
// 	}

// 	var payload jsonResponse
// 	payload.Error = false
// 	payload.Message = "Authenticated!"
// 	payload.Data = jsonFromService.Data

// 	app.writeJSON(w, http.StatusAccepted, payload)
// }

// // sendMail sends email by calling the mail microservice
// func (app *Config) sendMail(w http.ResponseWriter, msg MailPayload) {
// 	jsonData, _ := json.MarshalIndent(msg, "", "\t")

// 	// call the mail service
// 	mailServiceURL := "http://mailer-service/send"

// 	// post to mail service
// 	request, err := http.NewRequest("POST", mailServiceURL, bytes.NewBuffer(jsonData))
// 	if err != nil {
// 		app.errorJSON(w, err)
// 		return
// 	}

// 	request.Header.Set("Content-Type", "application/json")

// 	client := &http.Client{}
// 	response, err := client.Do(request)
// 	if err != nil {
// 		app.errorJSON(w, err)
// 		return
// 	}
// 	defer response.Body.Close()

// 	// make sure we get back the right status code
// 	if response.StatusCode != http.StatusAccepted {
// 		app.errorJSON(w, errors.New("error calling mail service"))
// 		return
// 	}

// 	// send back json
// 	var payload jsonResponse
// 	payload.Error = false
// 	payload.Message = "Message sent to " + msg.To

// 	app.writeJSON(w, http.StatusAccepted, payload)

// }

// // logEventViaRabbit logs an event using the logger-service. It makes the call by pushing the data to RabbitMQ.
// func (app *Config) logEventViaRabbit(w http.ResponseWriter, l LogPayload) {
// 	err := app.pushToQueue(l.Name, l.Data)
// 	if err != nil {
// 		app.errorJSON(w, err)
// 		return
// 	}

// 	var payload jsonResponse
// 	payload.Error = false
// 	payload.Message = "logged via RabbitMQ"

// 	app.writeJSON(w, http.StatusAccepted, payload)
// }

// // pushToQueue pushes a message into RabbitMQ
// func (app *Config) pushToQueue(name, msg string) error {
// 	emitter, err := event.NewEventEmitter(app.Rabbit)
// 	if err != nil {
// 		return err
// 	}

// 	payload := LogPayload{
// 		Name: name,
// 		Data: msg,
// 	}

// 	j, err := json.MarshalIndent(&payload, "", "\t")
// 	if err != nil {
// 		return err
// 	}

// 	err = emitter.Push(string(j), "log.INFO")
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

type RPCPayload struct {
	Name string
	Data string
}

// logItemViaRPC logs an item by making an RPC call to the logger microservice
// func (app *Config) logItemViaRPC(w http.ResponseWriter, l LogPayload) {
// 	client, err := rpc.Dial("tcp", "logger-service:5001")
// 	if err != nil {
// 		app.errorJSON(w, err)
// 		return
// 	}

// 	rpcPayload := RPCPayload{
// 		Name: l.Name,
// 		Data: l.Data,
// 	}

// 	var result string
// 	err = client.Call("RPCServer.LogInfo", rpcPayload, &result)
// 	if err != nil {
// 		app.errorJSON(w, err)
// 		return
// 	}

// 	payload := jsonResponse{
// 		Error:   false,
// 		Message: result,
// 	}

// 	app.writeJSON(w, http.StatusAccepted, payload)
// }

// func (app *Config) LogViaGRPC(w http.ResponseWriter, r *http.Request) {
// 	var requestPayload RequestPayload

// 	err := app.readJSON(w, r, &requestPayload)
// 	if err != nil {
// 		app.errorJSON(w, err)
// 		return
// 	}

// 	conn, err := grpc.Dial("logger-service:50001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
// 	if err != nil {
// 		app.errorJSON(w, err)
// 		return
// 	}
// 	defer conn.Close()

// 	c := logs.NewLogServiceClient(conn)
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
// 	defer cancel()

// 	_, err = c.WriteLog(ctx, &logs.LogRequest{
// 		LogEntry: &logs.Log{
// 			Name: requestPayload.Log.Name,
// 			Data: requestPayload.Log.Data,
// 		},
// 	})
// 	if err != nil {
// 		app.errorJSON(w, err)
// 		return
// 	}

// 	var payload jsonResponse
// 	payload.Error = false
// 	payload.Message = "logged"

// 	app.writeJSON(w, http.StatusAccepted, payload)
// }