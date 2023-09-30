package server

import (
	"context"
	"fmt"
	"log"

	greetingpb "github.com/ptrenwith/go_api/proto/pb/greeting"
)

type GreetingService struct {
	greetingpb.UnimplementedGreeterServer
}

func (h GreetingService) SayHello(ctx context.Context, req *greetingpb.HelloRequest) (*greetingpb.HelloReply, error) {
	greeting := fmt.Sprintf("Hello %s", req.Name)
	log.Println(greeting)
	return &greetingpb.HelloReply{
		Message: greeting,
	}, nil
}
