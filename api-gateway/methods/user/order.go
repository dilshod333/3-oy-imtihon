package user1

import (
	"api-gateway/proto/gen1"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func OrderProductt() gen1.OrderServiceClient{
	conn, err := grpc.NewClient(":9999", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal("errror oonnn clientcalll...", err)
	}
	client := gen1.NewOrderServiceClient(conn)

	return client
}
