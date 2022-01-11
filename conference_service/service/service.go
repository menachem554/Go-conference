package tickets

import (
	"context"
	"fmt"
	"log"

	pb "github.com/menachem554/Go-conference/conference_service/proto"
	dbModel "github.com/menachem554/Go-conference/conference_service/service/mongo"
	"github.com/spf13/viper"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)


type server struct {
	conferenceModelController dbModel.ConferenceMongoController
	pb.UnimplementedGoConferenceServer
}

func CreateConferenceService() pb.GoConferenceServer {
	mongoController, err := dbModel.CreateMongoUsersController(viper.GetString("MongoConnString"), viper.GetString("conferenceCollection"))
	if err != nil {
		log.Fatal("could not create the service")
	}

	return &server{conferenceModelController: mongoController}
}

func OrderToProto(data *dbModel.OrderTicketModel) *pb.UserData {
	return &pb.UserData{
		FirstName:   	data.FirstName,
		LastName:		data.LastName,
		Email: 			data.Email,
		NumberOfTicket: data.NumberOfTicket,
	}
}
func (s *server) PostNewOrder(ctx context.Context, order *pb.UserDataReq) (*pb.UserDataRes, error) {
	// Get the request from grpc
	req := order.GetUserData()

	fmt.Println(req)
	fmt.Printf("the type: %v\n", req)

	data := dbModel.OrderTicketModel {
		FirstName: req.FirstName,
		LastName: req.LastName,
		Email: req.Email,
		NumberOfTicket: req.NumberOfTicket,
	}
	fmt.Println(data)
	fmt.Printf("the type: %v\n", data)

	res, err := s.conferenceModelController.CreateNewOrder(ctx, data)
	if err != nil {
		return nil,
		status.Errorf(codes.Internal, fmt.Sprintf("Internal Error: %v", err))
	}

	fmt.Printf("Successfully inserted the new order into orders collection!, %v \n", res)

	return &pb.UserDataRes{UserData: OrderToProto(&data)}, nil
}

func(s *server) GetOrder(ctx context.Context, email *pb.GetOrderReq) (*pb.UserDataRes, error) {
	// Get the request from grpc
	// req := email

	// data := dbModel.OrderTicketModel{Email: req}
	// fmt.Printf("the email you write is: %T\n", data)
	fmt.Println("email at grpc req:", email.GetEmail())

	res, err := s.conferenceModelController.GetOrder(ctx, email.GetEmail())
	if err != nil {
		return nil,
		status.Errorf(codes.Internal, fmt.Sprintf(" Internal Error: %v", err))
	}

	return &pb.UserDataRes{UserData: OrderToProto(res)}, nil
}