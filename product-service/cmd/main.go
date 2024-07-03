package main

import (
	"log"
	"net"
	"product-service/api"
	"product-service/internal/proto/gen3"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    } 
    grpcServer := grpc.NewServer()
	
	server2 := api.DbConn()
    gen3.RegisterProductServiceServer(grpcServer, server2)
	
	reflection.Register(grpcServer)
	
    log.Println("Server running on :8080")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
