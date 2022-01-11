package mongo

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"github.com/menachem554/Go-conference/conference_service/database"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type ConferenceMongoController interface {
	CreateNewOrder(ctx context.Context, user OrderTicketModel) (*OrderTicketModel, error)
	GetOrder(ctx context.Context, email string) (*OrderTicketModel, error)
	UpdateOrder(ctx context.Context, email string, numOfTickets int32) (*OrderTicketModel, error)
}

type conferenceMongoController struct {
	mongoClient *mongo.Client
	collection  *mongo.Collection
}

func CreateMongoUsersController(mongoDBConnectionString string, usersCollection string) (ConferenceMongoController, error) {
	mongoClient, err := database.ConnectToMongo(mongoDBConnectionString)
	if err != nil {
		return nil, err
	}

	mongoDatabase, err := database.GetMongoDatabase(mongoClient, mongoDBConnectionString)
	if err != nil {
		return nil, err
	}

	collection := mongoDatabase.Collection(usersCollection)

	
	return &conferenceMongoController{mongoClient, collection}, nil
}

func(cc *conferenceMongoController) CreateNewOrder(ctx context.Context, newOrder OrderTicketModel) (*OrderTicketModel, error) {
	data, err := cc.collection.InsertOne(ctx, newOrder)

	if err !=  nil  {
		return nil, status.Error(codes.Internal, err.Error())
	}

	fmt.Printf("new order is good the objectID is: %v\n", data.InsertedID)

	return &newOrder, nil 
}

func(cc *conferenceMongoController) GetOrder(ctx context.Context, email string) (*OrderTicketModel, error) {
	// search order by the email that the user enter
	res := cc.collection.FindOne(ctx, bson.M{"email": email})

	err := res.Err()
	if err == mongo.ErrNoDocuments {
		return nil, status.Error(codes.NotFound, "Cannot found an order with that email")
	}

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	data := OrderTicketModel{}

	// decode and Check for error
	if err := res.Decode(&data); err != nil {
		return nil,
			status.Errorf(codes.NotFound, fmt.Sprintf("Cannot decode the response from mongo: %v", err))
	}

	fmt.Println("the addres response from the server is: ", &res)
	return &data, nil
}

func(cc *conferenceMongoController) UpdateOrder(ctx context.Context, email string, numOfTickets int32) (*OrderTicketModel, error) {
	// insert the changes
	res :=  cc.collection.FindOneAndUpdate(
		ctx,
		bson.M{"email":email},
		bson.M{"$set": numOfTickets})

		err := res.Err()
	if err == mongo.ErrNoDocuments {
		return nil, status.Error(codes.NotFound, "Cannot found an order with that email")
	}

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return nil, nil
}
