package handleproduct

import (
	"context"
	"log"
	"order-service/api/models"
	"order-service/proto/gen3"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// import (
// 	"context"
// 	"log"
// 	"order-service/api/models"
// 	"order-service/proto/product"

// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/credentials/insecure"
// )

// func HandleClient(req int) *models.Handle {
// 	conn, err := grpc.NewClient(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	client := product.NewProductServiceClient(conn)

// 	res, err := client.GetProducts(context.Background(), &product.ProductId{ProductId: int32(req)})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return &models.Handle{
// 		NameProduct: res.Name,
// 		Price:       res.Price,
// 		Type:        res.Type,
// 		Quantity:    int(res.Quantity),
// 	}
// }


func ProductHandle(req int) *models.Handle{
	conn, err := grpc.NewClient(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal("Something wrongg...", err)
	}
	client := gen3.NewProductServiceClient(conn)

	res, err := client.GetProducts(context.Background(), &gen3.ProductId{ProductId: int32(req)})

	if err != nil {
		log.Fatal("getproduct oganda xatoo....", err)
	}
	
	return &models.Handle{
		NameProduct: res.Name,
		Price:       res.Price,
		Type:        res.Type,
		Quantity:    int(res.Quantity),
	}
	
}