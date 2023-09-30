package server

// go mod init github.com/ptrenwith/go_api/server
import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	greetingpb "github.com/ptrenwith/go_api/proto/pb/greeting"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	grpcPort        = 50051
	grpcGatewayPort = 50052
)

func StartServer() {
	go StartGRPCServer()
	StartGRPCGateway()
}

func StartGRPCServer() {
	log.Printf("Starting gRPC Server...")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("gRPC Server failed to listen: %v", err)
	}

	// Setup gRPC Service
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	// Register services
	greetingpb.RegisterGreeterServer(grpcServer, GreetingService{})

	// Start Server
	log.Printf("gRPC Server listening on: %s", lis.Addr())
	grpcServer.Serve(lis)
}

func StartGRPCGateway() {
	log.Printf("Starting gRPC Gateway...")

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		fmt.Sprintf("0.0.0.0:%d", grpcPort),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial gRPC server:", err)
	}

	gwmux := runtime.NewServeMux()
	// Register Greeter
	err = greetingpb.RegisterGreeterHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", grpcGatewayPort),
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on", gwServer.Addr)
	log.Fatalln(gwServer.ListenAndServe())
}
