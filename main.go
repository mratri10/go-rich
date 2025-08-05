package main

import (
	"log"
	"net"

	_ "github.com/lib/pq"
	"github.com/mratri10/go-rich/db"
	"github.com/mratri10/go-rich/pb/github.com/atri/go-grpc-purchase/pb"
	"github.com/mratri10/go-rich/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	dbConn, err := db.NewDB("postgres://richman:glory_100jt@167.99.76.27:5432/richdb?sslmode=disable")

	if err != nil {
		log.Fatal("Database connection failed: ", err)
	}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("Failed to listen:", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterPurchaseServiceServer(grpcServer, &server.PurchaseServer{DB: dbConn})

	reflection.Register(grpcServer)

	log.Println("gRPC server started on port 50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to serve: ", err)
	}
}
