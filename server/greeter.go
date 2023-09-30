package server

import (
	"context"
	"fmt"

	pb "github.com/ptrenwith/go_api/messages/pb"
)

type GreetingService struct {
	pb.UnimplementedGreeterServer
}

func (h GreetingService) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	greeting := fmt.Sprintf("Hello %s", req.Name)
	return &pb.HelloReply{
		Message: greeting,
	}, nil
}
