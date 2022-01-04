package server

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	pb "github.com/menachem554/Go-conference/"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGoConferenceServer
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


func ConnectToGrpc() {
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
	fmt.Println("Closing MongoDB...")
	client.Disconnect(mongoCtx)
	fmt.Println("All done!")
}