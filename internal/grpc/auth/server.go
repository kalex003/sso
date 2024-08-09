package auth

import (
	"context"
	ssov1 "github.com/kalex003/protos/gen/go/sso"
	"google.golang.org/grpc"
)

type serverAPI struct { //для реализация интерфейсов, сгенерированнхы прото файлом
	ssov1.UnimplementedLoginServiceServer //метод, сгенерированный grpc
}

func Register(gRPC *grpc.Server) {
	ssov1.RegisterAuthServer(gRPC, &serverAPI{})
}

func (s *serverAPI) Login(ctx context.Context, req *ssov1.LoginRequest) (*ssov1.LoginResponse, error) {
	panic("implement me")
}

func (s *serverAPI) Register(ctx context.Context, req *ssov1.LoginRequest) (*ssov1.LoginResponse, error) {
	panic("implement me")
}

func (s *serverAPI) IsAdmin(ctx context.Context, req *ssov1.LoginRequest) (*ssov1.LoginResponse, error) {
	panic("implement me")
}
