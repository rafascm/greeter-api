package svc

import (
	"context"
	"fmt"

	pb "github.com/rafascm/greeter-api/pkg/pb/greeting/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GreeterService struct todos os handlers impostos
// na definição do nosso serviço gRPC.
type GreeterService struct{}

// Greet retorna uma mensagem concatenada com o cumprimento e o nome
// da pessoa que recebemos na requisição.
func (gs GreeterService) Greet(ctx context.Context, req *pb.GreetRequest) (res *pb.GreetResponse, err error) {
	if req.Msg == nil {
		err = status.New(codes.InvalidArgument, "Name cannot be empty").Err()
		return
	}

	helloMsg := fmt.Sprintf("Hello, %s!", req.Msg.Name)

	res = &pb.GreetResponse{
		Resp: helloMsg,
	}

	return
}
