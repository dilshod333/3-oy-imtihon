package main

import (
	"log"
	"net"
	"user-service/api"
	"user-service/config"
	"user-service/internal/proto/gen"
	"user-service/logs"
	"user-service/pkg"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lg := logs.GetLogger("./logs/logger.log")

	cfg := config.Load(".env")
	connDb, err := pkg.ConnectToDB(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer connDb.Close()

	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()
	grpcServer := grpc.NewServer()

	server, err := api.NewServer(lg)
	if err != nil {
		log.Fatal(err)
	}
	gen.RegisterUserServiceServer(grpcServer, server)

	reflection.Register(grpcServer)
	lg.Println("Mainda ishlatdim loggerni...")

	log.Println("Server running on :8082")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
