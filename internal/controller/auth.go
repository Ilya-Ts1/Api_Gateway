package controller

import (
	"fmt"
	"time"

	pb "api_gateway/protos/gen/auth"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

type AuthClient struct {
	conn   *grpc.ClientConn
	Client pb.AuthServiceClient
}

func NewAuthClient(addr string) (*AuthClient, error) {
	conn, err := grpc.Dial(addr,
		grpc.WithInsecure(),
		grpc.WithTimeout(5*time.Second),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                30 * time.Second,
			Timeout:             5 * time.Second,
			PermitWithoutStream: true,
		}),
	)
	if err != nil {
		return nil, fmt.Errorf("grpc.Dial %s: %w", addr, err)
	}

	return &AuthClient{
		conn:   conn,
		Client: pb.NewAuthServiceClient(conn),
	}, nil
}
