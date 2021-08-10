package service

import (
	"context"
	"fmt"

	"github.com/its-dastan/grpcDemo/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthServer struct {
	pb.UnimplementedAuthServiceServer
	userStore  UserStore
	jwtManager *JWTManager
}

func NewAuthService(userStore UserStore, jwtManager *JWTManager) *AuthServer {
	return &AuthServer{userStore: userStore, jwtManager: jwtManager}
}

func (server *AuthServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := server.userStore.Find(req.GetUsername())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot find user : %v", err)
	}

	if user == nil {
		return nil, status.Errorf(codes.NotFound, "incorrect username/password")
	}

	token, err := server.jwtManager.Generate(user)
	fmt.Println(err)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token : %v", err)
	}
	res := &pb.LoginResponse{AccessToken: token}

	return res, nil
}
