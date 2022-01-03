package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	// "time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	// "github.com/joho/godotenv"
	pb "github.com/menachem554/Go-conference/backend/proto"
	repo "github.com/menachem554/Go-conference/backend/src/conference.mongo.go"
)

type server struct {
	pb.UnimplementedGoConferenceServer
}

// mongo setting
// var db *mongo.Client
var ticketDB *mongo.Collection
var mongoCtx context.Context

// UserData interface
type DataToMongo struct {
	FirstName      string `bson:firstName`
	LastName       string `bson:lastName`
	Email          string `bson:email`
	NumberOfTicket uint32 `bson:numberOfTicket`
}

// UserData to pb response
func DataToGrpc(data *DataToMongo) *pb.UserData {
	return &pb.UserData{
		FirstName:      data.FirstName,
		LastName:       data.LastName,
		Email:          data.Email,
		NumberOfTicket: data.NumberOfTicket,
	}
}

// Create new Order
func (s *server) PostNewOrder(ctx context.Context, req *pb.UserDataReq) (*pb.UserDataRes, error) {
	// Get the request
userReq := req.GetUserData()

	data := DataToMongo{
		FirstName: userReq.GetFirstName(),
		LastName: userReq.GetLastName(),
		Email: userReq.GetEmail(),
		NumberOfTicket: userReq.GetNumberOfTicket(),
	}

	// InsertNewOrder()

	// Insert the data into the database
	// res, err := ticketDB.InsertOne(mongoCtx, data)
	// if err != nil {
	// 	return nil,
	// 		status.Errorf(codes.Internal, fmt.Sprintf(" Internal Error: %v", err))
	// }

	// fmt.Printf("Successfully inserted new order into orders collection!, %v \n", res)

	return &pb.UserDataRes{UserData: DataToGrpc(&data)}, nil
}

// Get order by the email of the order
func (s *server) GetOrder(ctx context.Context, req *pb.GetOrderReq) (*pb.UserDataRes, error) {
	// Get the email 
	userEmail := req.GetEmail()

	res := ticketDB.FindOne(ctx, bson.M{"email": userEmail})
	data := &DataToMongo{}

	// decode and Check for error
	if err := res.Decode(data); err != nil {
		return nil,
			status.Errorf(codes.NotFound, fmt.Sprintf("Cannot found user with this Email: %v", err))
	}
	fmt.Println("Get user result", data)
	return &pb.UserDataRes{UserData: DataToGrpc(data)}, nil
}

// Update the ticket for the user
func (s *server) UpdateOrder(ctx context.Context, req *pb.UpdateOrderRequest) (*pb.UserDataRes, error) {
	// Get the request
	userEmail := req.GetEmail()
	ticketNumber := req.GetNumberOfTicketd()

	// insert the new number of ticket
	ticketDB.FindOneAndUpdate(
		ctx,
		bson.M{"email": userEmail},
		bson.D{
			{Key: "$set", Value: bson.D{
				primitive.E{Key: "numberOfTicket", Value: ticketNumber},
			}},
		},
	)

	// Get the update number rof ticket
	res := ticketDB.FindOne(ctx, bson.M{"email": userEmail})
	data := &DataToMongo{}

	// decode and Check for error
	if err := res.Decode(data); err != nil {
		return nil,
			status.Errorf(codes.NotFound, fmt.Sprintf("Cannot found user with this Email: %v", err))
	}
	fmt.Println("Get user result", data)
	return &pb.UserDataRes{UserData: DataToGrpc(data)}, nil
}

// //
// func (s *server) DeleteBook(ctx context.Context, req *pb.GetBookReq) (*pb.DeleteBookRes, error) {
// 	// Get ID of the book
// 	bookID := req.GetId()

// 	res, err := bookDB.DeleteOne(ctx, bson.M{"bookid": bookID})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Book deleted: ", res.DeletedCount)

// 	return &pb.DeleteBookRes{Deleted: res.DeletedCount}, nil
// }

// // Get all books in the Collection
// func (s *server) GetAllBooks(ctx context.Context, req *pb.GetAllReq) (*pb.GetAllResponse, error) {
// 	fmt.Println("\n list of all book start stream")

// 	res, err := bookDB.Find(context.Background(), bson.M{})
// 	if err != nil {
// 		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Unknown Internal Error: %v", err))
// 	}

// 	defer res.Close(context.Background())

// 	var books = []*BookInterface{}

// 	for res.Next(context.Background()) {
// 		var data = &BookInterface{}
// 		if err := res.Decode(data); err != nil {
// 			return nil, status.Errorf(codes.Internal, fmt.Sprintf("Cannot decoding data: %v", err))
// 		}
// 		books = append(books, data)
// 	}
// 	if err = res.Err(); err != nil {
// 		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Unknown Internal Error: %v", err))
// 	}
// 	var pbbooks = []*pb.Book{}
// 	for _, data := range books {
// 		fmt.Println(data)
// 		pbbooks = append(pbbooks, &pb.Book{BookID: data.BookID, BookName: data.BookName,
// 			Category: data.Category, Author: data.Author})
// 	}
// 	fmt.Println(pbbooks)
// 	return &pb.GetAllResponse{Book: pbbooks}, nil
// }

func main() {
	// log if go crash, with the file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	repo.ConnectToMongo()

	// // get env vars
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatalf("Error loading .env file")
	// }

	// mongoLocal := os.Getenv("MONGO_LOCAL")
	// // mongoImage := os.Getenv("MONGO_IMAGE")

	// // create the mongo context
	// mongoCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	// // connect MongoDB d
	// fmt.Println("Connecting to MongoDB...")
	// client, err := mongo.Connect(mongoCtx, options.Client().ApplyURI(mongoLocal))
	// if err != nil {
	// 	log.Fatalf("Error Starting MongoDB Client: %v", err)
	// }

	// // check the connection
	// err = client.Ping(mongoCtx, nil)
	// if err != nil {
	// 	log.Fatalf("Could not connect to MongoDB: %v\n", err)
	// } else {
	// 	fmt.Println("Connected to Mongodb")
	// }

	// ticketDB = client.Database("GoConference").Collection("orders")

	fmt.Println("Starting Listener...")
	l, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	pb.RegisterGoConferenceServer(s, &server{})

	// Start a GO Routine
	go func() {
		fmt.Println(" Server GoConference Started...")
		if err := s.Serve(l); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait to exit (Ctrl+C)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Block the channel until the signal is received
	<-ch
	fmt.Println("Stopping GoConference Server...")
	s.Stop()
	fmt.Println("Closing Listener...")
	l.Close()
	// fmt.Println("Closing MongoDB...")
	// client.Disconnect(mongoCtx)
	fmt.Println("All done!")
}
