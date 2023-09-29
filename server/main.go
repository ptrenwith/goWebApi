package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/ptrenwith/goWebApi/greetings/pb"
	"google.golang.org/grpc"
)

const (
	port = 50051
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	grpcServer.Serve(lis)
}

func GreetVisitor(req *pb.HelloRequest) *pb.HelloReply {
	greeting := fmt.Sprintf("Hello %s", req.Name)
	return &pb.HelloReply{
		Message: greeting,
	}
}
