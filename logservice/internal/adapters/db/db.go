package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ayuved/microservices-helper/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Logservice struct {
	ID   string
	App  string
	Name string
	Data string
}

type Adapter struct {
	db *mongo.Collection
}

func (a Adapter) Get(ctx context.Context, id string) (domain.Logservice, error) {

	// Find
	var lsModel Logservice
	err := a.db.FindOne(ctx, bson.M{"id": id}).Decode(&lsModel)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found document: %+v\n", lsModel)

	result := domain.Logservice{
		ID:   lsModel.ID,
		App:  lsModel.App,
		Name: lsModel.Name,
		Data: lsModel.Data,
	}
	return result, err
}

func (a Adapter) Add(ctx context.Context, logservice *domain.Logservice) error {
	log.Println("Inserting logservice", logservice)
	logModel := Logservice{

		App:  logservice.App,
		Name: logservice.Name,
		Data: logservice.Data,
	}
	log.Println("Inserting logservice", logModel)
	insertResult, err := a.db.InsertOne(ctx, logModel)
	log.Println("Inserting logservice", insertResult)
	log.Println("Inserting logservice", err)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Inserted document with ID:", insertResult.InsertedID)
	// fmt.Printf("Inserted document with ID: %v\n", insertResult.InsertedID)

	return err
}

func NewAdapter(dataSourceUrl string) (*Adapter, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dataSourceUrl))
	if err != nil {
		log.Fatal(err)
	}
	// defer client.Disconnect(ctx)
	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("could not ping MongoDB: %v", err)
	}

	// Get collection
	db := client.Database("logservice").Collection("logs")

	log.Println("Connecting6 to MongoDB...", db)
	//err := db.AutoMigrate(&Logservice{})
	//if err != nil {
	//	return nil, fmt.Errorf("db migration error: %v", err)
	//}
	return &Adapter{db: db}, nil
}
