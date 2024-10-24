package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ayuved/microservices-helper/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Logservice struct {
	ID        string
	App       string
	Name      string
	Type      string
	Status    string
	ProcessId string
	Data      interface{}
	User      string
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
	DeletedAt time.Time
	DeletedBy string
}

type Adapter struct {
	db *mongo.Collection
}

func (a Adapter) Get(ctx context.Context, id string) (domain.Logservice, error) {

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.Logservice{}, status.Errorf(codes.InvalidArgument, "Cannot parse ID")
	}
	// Find
	var lsModel Logservice
	filter := bson.M{"_id": oid}
	err = a.db.FindOne(ctx, filter).Decode(&lsModel)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found document: %+v\n", lsModel)

	result := domain.Logservice{
		ID:        lsModel.ID,
		App:       lsModel.App,
		Name:      lsModel.Name,
		Type:      lsModel.Type,
		Status:    lsModel.Status,
		ProcessId: lsModel.ProcessId,
		Data:      lsModel.Data.(string),
		User:      lsModel.User,
	}
	return result, err
}

func (a Adapter) Add(ctx context.Context, logservice *domain.Logservice) (error, string)  {
	log.Println("Inserting logservice", logservice)
	logModel := Logservice{
		App:       logservice.App,
		Name:      logservice.Name,
		Type:      logservice.Type,
		Status:    logservice.Status,
		ProcessId: logservice.ProcessId,
		Data:      logservice.Data,
		CreatedAt: time.Now(),
		CreatedBy: logservice.User,
	}
	log.Println("Inserting logservice", logModel)
	insertResult, err := a.db.InsertOne(ctx, logModel)
	if err != nil {
		log.Fatal(err)
	}
	oid, ok := insertResult.InsertedID.(primitive.ObjectID)
	if !ok {
		return status.Errorf(codes.Internal, "Cannot convert to OID"), ""
	}
	log.Println("Inserted document with ID:", oid.Hex())
	// fmt.Printf("Inserted document with ID: %v\n", insertResult.InsertedID)

	return err, oid.Hex()
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
