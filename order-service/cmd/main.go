package main

import (
	"log"
	"net"
	"order-service/api/handlers"
	"order-service/logs"
	"order-service/pkg"
	"order-service/proto/gen1"
	"os"
	
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	_"github.com/joho/godotenv/autoload"
	_"github.com/lib/pq"
)

func main() {

	lgs := logs.GetLogger("./logs/logger.log")

	// cfg := config.Load(".env")
	DB := pkg.OpenSql(os.Getenv("DATABASE_URL"))
	// connDb, err := pkg.ConnectToDB(cfg)
	// if err != nil {
	// 	fmt.Println("Failed to connect to the database:", err)
	// 	return
	// }
	// defer connDb.Close()

	lis, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	lgs.Println("Eshityaptii...:9999")
	grpcServer := grpc.NewServer()

	server := handlers.NewUserServer(DB, lgs)
	if err != nil {
		log.Fatal(err)
	}
	gen1.RegisterOrderServiceServer(grpcServer, server)

	reflection.Register(grpcServer)
	lgs.Println("Mainda ishlatdim loggerni...")

	log.Println("Server running on :9999")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
