package handleuser

import (
	"context"
	"log"
	"order-service/proto/gen"
	// "order-service/proto/gen1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func UserClinet(req int) (string, error) {
	conn, err := grpc.NewClient(":8082", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	client := gen.NewUserServiceClient(conn)

	res, err := client.GetUser(context.Background(), &gen.GetUserReq{UserId: int32(req)})
	if err != nil {
		log.Fatal(err)
	}

	return res.Name, nil 
}

