package database

import (
	"context"
	"fmt"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

func ConnectToMongo(ConnectionString string) (*mongo.Client, error) {
	// Create client
	mongoOptions := options.Client().ApplyURI(ConnectionString)
	client, err := mongo.NewClient(mongoOptions)
	if err != nil {
		return nil, fmt.Errorf("failed creating mongodb client with connection string \"%s\": %v", ConnectionString, err)
	}

	// Connect client to mongo
	connTimeoutCtx, cancelConn := context.WithTimeout(context.Background(), viper.GetDuration("MongoClientConnectionTimout"))
	defer cancelConn()
	err = client.Connect(connTimeoutCtx)
	if err != nil {
		return nil, fmt.Errorf("failed connecting to mongodb with connection string %s: %v", ConnectionString, err)
	}

	// Check the connection.
	pingTimeoutCtx, cancelPing := context.WithTimeout(context.Background(), viper.GetDuration("MongoClientPingTimeout"))
	defer cancelPing()
	err = client.Ping(pingTimeoutCtx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed pinging to mongodb with connection string %s: %v", ConnectionString, err)
	}

	return client, nil
}

func GetMongoDatabase(client *mongo.Client, connString string) (*mongo.Database, error) {
	ConnString, err := connstring.ParseAndValidate(connString)
	if err != nil {
		return nil, fmt.Errorf("failed parsing connection string %s: %v", connString, err)
	}

	return client.Database(ConnString.Database), nil
}

// 	// Check for error while load the env var
// 	err := godotenv.Load(".env")
// 	if err != nil {
// 		log.Fatalf("Error loading .env file")
// 	}

// 	// Choose the url for mongo,  localhost || image
// 	isLocal:= true
// 	var mongoAddress string
// 	if isLocal {
// 		mongoAddress = os.Getenv("MONGO_LOCAL")
// 	} else {
// 		mongoAddress = os.Getenv("MONGO_IMAGE")
// 	}

// 	// create the mongo context
// 	mongoCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	// connect MongoDB
// 	fmt.Println("Connecting to MongoDB...")
// 	client, err := mongo.Connect(mongoCtx, options.Client().ApplyURI(mongoAddress))
// 	if err != nil {
// 		log.Fatalf("Error Starting MongoDB Client: %v", err)
// 	}

// 	// check the connection
// 	err = client.Ping(mongoCtx, nil)
// 	if err != nil {
// 		log.Fatalf("Could not connect to MongoDB: %v\n", err)
// 	} else {
// 		fmt.Println("Connected to Mongodb")
// 	}

// 	client.Database("Bookstore").Collection("books")

// 	return client, nil
// }
