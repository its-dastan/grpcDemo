package client

import (
	"context"
	"time"

	"github.com/its-dastan/grpcDemo/pb"
	"google.golang.org/grpc"
)

type AuthClient struct {
	service  pb.AuthServiceClient
	username string
	password string
}

func NewAuthClient(cc *grpc.ClientConn, username string, password string) *AuthClient {
	service := pb.NewAuthServiceClient(cc)
	return &AuthClient{
		service, username, password,
	}
}

func (client *AuthClient) Login() (string, error) {
	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()

	req := &pb.LoginRequest{
		Username: client.username,
		Password: client.password,
	}

	res, err := client.service.Login(ctx, req)
	if err != nil {
		return "", err
	}
	return res.GetAccessToken(), nil
}