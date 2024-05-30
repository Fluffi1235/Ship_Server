package auth

import (
	"context"
	ssov1 "diplomServer/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ServerAuthAPI struct {
	ssov1.UnimplementedAuthServer
	Auth *AuthImpl
}

func Register(gRPC *grpc.Server, auth *AuthImpl) {
	ssov1.RegisterAuthServer(
		gRPC,
		&ServerAuthAPI{Auth: auth},
	)
}

func (s *ServerAuthAPI) Register(ctx context.Context, req *ssov1.RegisterRequest) (*ssov1.RegisterResponse, error) {
	if req.UserName == "" {

		return nil, status.Error(codes.InvalidArgument, "username is required")
	}

	if req.Password == nil {

		return nil, status.Error(codes.InvalidArgument, "password is required")
	}

	resp, err := s.Auth.RegisterNewUser(ctx, req)
	if err != nil {

		return nil, status.Error(codes.Internal, "internal error")
	}

	return resp, nil
}

func (s *ServerAuthAPI) Login(ctx context.Context, req *ssov1.LoginRequest) (*ssov1.LoginResponse, error) {
	if req.UserName == "" {
		return nil, status.Error(codes.InvalidArgument, "username is required")
	}

	if req.Password == nil {
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}

	resp, err := s.Auth.Login(ctx, req)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	return resp, nil
}
