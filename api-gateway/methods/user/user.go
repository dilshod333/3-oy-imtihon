package user1

import (
	"api-gateway/proto/gen"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ConnectUser() gen.UserServiceClient{
	conn, err := grpc.NewClient(":8082", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal("error connect user micro...", err)
	}

	client := gen.NewUserServiceClient(conn)
	return client 
}