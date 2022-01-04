package mongo

import (


	"go.mongodb.org/mongo-driver/mongo"
	server "github.com/menachem554/Go-conference/conference_service/server"
)

type ticketsMongoController struct {
	mongoClient *mongo.Client
	collection *mongo.Collection
}

type DataToMongo struct {

}

func CreateNewOrder(ctx context.Context userReq  ) (error) {
	data, err := ticketDB.InsertOne(mongoCtx, userReq)
	if err != nil {
		return nil,
			status.Errorf(codes.Internal, fmt.Sprintf(" Internal Error: %v", err))
	}

	fmt.Printf("Successfully inserted NEW Book into book collection!, %v \n", res)


	return data, nil
}

ConnectToMongo()
server