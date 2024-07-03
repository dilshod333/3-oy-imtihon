package user1

import (
	"api-gateway/proto/gen3"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ConnectProduct() gen3.ProductServiceClient{
	conn, err := grpc.NewClient(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal("errror oonnn clientcalll...", err)
	}
	client := gen3.NewProductServiceClient(conn)

	return client
}
