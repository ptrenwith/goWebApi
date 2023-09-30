package server

// go mod init github.com/ptrenwith/go_api/server
import (
	"fmt"
	"log"
	"net"

	pb "github.com/ptrenwith/go_api/messages/pb"
	"google.golang.org/grpc"
)

const (
	port = 50051
)

func StartServer() {
	log.Printf("Server starting...")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterGreeterServer(grpcServer, GreetingService{})
	log.Printf("Server listening on: %s", lis.Addr())
	grpcServer.Serve(lis)
}
