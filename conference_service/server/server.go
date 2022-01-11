package server

import (
	"fmt"
	"net"

	pb "github.com/menachem554/Go-conference/conference_service/proto"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

type ConferenceServer struct {
	grpcServer *grpc.Server
	protoServer *pb.GoConferenceServer
	listener   *net.Listener
	port       string
}

func NewServer(protoServer *pb.GoConferenceServer, port string) *ConferenceServer{
	s := &ConferenceServer{protoServer: protoServer, port: port}

	serverOpts := []grpc.ServerOption{grpc.MaxRecvMsgSize(viper.GetInt("GrpcMaxReceivedMessageSize"))}

	s.grpcServer = grpc.NewServer(serverOpts...)

	pb.RegisterGoConferenceServer(s.grpcServer, *protoServer)

	listener, err := net.Listen("tcp", ":"+port)

	if err != nil {
		fmt.Errorf("failed to create listener")
	}
	s.listener = &listener

	// fmt.Printf("Listening on port: %v\n", s.port)

	return s
}

func (uss *ConferenceServer) Serve() {
	fmt.Println("Listening on port:", uss.port)
	err := uss.grpcServer.Serve(*uss.listener)
	if err != nil {
		fmt.Printf("failed to serve%V\n", err)
	}
	// utils.ValidateOrPanic(err, "failed to serve", uss.logger)
}

// func StartGrpcServer() {
// 	// log if go crash, with the file name and line number
// 	log.SetFlags(log.LstdFlags | log.Lshortfile)

// 	fmt.Println("Starting Listener...")
// 	l, err := net.Listen("tcp", "0.0.0.0:9090")
// 	if err != nil {
// 		log.Fatalf("Failed to listen: %v", err)
// 	}

// 	opts := []grpc.ServerOption{}
// 	s := grpc.NewServer(opts...)
// 	pb.RegisterGoConferenceServer(s, pb.UnimplementedGoConferenceServer{})

// 	// Start a GO Routine
// 	go func() {
// 		fmt.Println("Bookstore Server Started...")
// 		if err := s.Serve(l); err != nil {
// 			log.Fatalf("Failed to start server: %v", err)
// 		}
// 	}()

// 	// Wait to exit (Ctrl+C)
// 	ch := make(chan os.Signal, 1)
// 	signal.Notify(ch, os.Interrupt)

// 	// Block the channel until the signal is received
// 	<-ch
// 	fmt.Println("Stopping Bookstore Server...")
// 	s.Stop()
// 	fmt.Println("Closing Listener...")
// 	l.Close()
// 	fmt.Println("Closing MongoDB...")
// 	fmt.Println("All done!")
// }