package backend

import (
	"context"

	pb "github.com/menachem554/Go-conference/backend/proto"
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

	InsertNewOrder(ctx, data)

	// Insert the data into the database
	// res, err := ticketDB.InsertOne(mongoCtx, data)
	// if err != nil {
	// 	return nil,
	// 		status.Errorf(codes.Internal, fmt.Sprintf(" Internal Error: %v", err))
	// }

	// fmt.Printf("Successfully inserted new order into orders collection!, %v \n", res)

	return &pb.UserDataRes{UserData: DataToGrpc(&data)}, nil
}