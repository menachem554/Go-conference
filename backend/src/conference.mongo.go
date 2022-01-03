package backend

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	// "github.com/menachem554/Go-conference/backend/main.go"
)


var ticketDB *mongo.Collection
var mongoCtx context.Context

type DataToMongo struct {
	FirstName      string `bson:firstName`
	LastName       string `bson:lastName`
	Email          string `bson:email`
	NumberOfTicket uint32 `bson:numberOfTicket`
}


func InsertNewOrder(ctx context.Context, data DataToMongo ) {
	// Insert the data into the database
	res, err := ticketDB.InsertOne(mongoCtx, data)
	if err != nil {
		status.Errorf(codes.Internal, fmt.Sprintf(" Internal Error: %v", err))
	}

	fmt.Printf("Successfully inserted new order into orders collection!, %v \n", res)
}

func ConnectToMongo() {
		// get env vars
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error loading .env file")
		}
	
		mongoLocal := os.Getenv("MONGO_LOCAL")
		// mongoImage := os.Getenv("MONGO_IMAGE")
	
		// create the mongo context
		mongoCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
	
		// connect MongoDB d
		fmt.Println("Connecting to MongoDB...")
		client, err := mongo.Connect(mongoCtx, options.Client().ApplyURI(mongoLocal))
		if err != nil {
			log.Fatalf("Error Starting MongoDB Client: %v", err)
		}
	
		// check the connection
		err = client.Ping(mongoCtx, nil)
		if err != nil {
			log.Fatalf("Could not connect to MongoDB: %v\n", err)
		} else {
			fmt.Println("Connected to Mongodb")
		}
	
		ticketDB = client.Database("GoConference").Collection("orders")


		// Wait to exit (Ctrl+C)
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, os.Interrupt)

		// Block the channel until the signal is received
		<-ch
		
		fmt.Println("Closing MongoDB...")
		client.Disconnect(mongoCtx)
}