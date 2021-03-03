package svc

import (
	"context"
	"fmt"

	pb "github.com/rafascm/greeter-api/pkg/pb/greeting/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GreeterService struct
type GreeterService struct{}

// Greet returns "hello, $name", name received from request
func (gs GreeterService) Greet(ctx context.Context, req *pb.GreetRequest) (res *pb.GreetResponse, err error) {
	if req.Name == "" {
		err = status.New(codes.InvalidArgument, "Name cannot be empty").Err()
		return
	}

	helloMsg := fmt.Sprintf("Hello, %s!", req.Name)

	res = &pb.GreetResponse{
		Resp: helloMsg,
	}

	return
}
