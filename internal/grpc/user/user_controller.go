package user

import (
	"context"
	ssov1 "diplomServer/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ServerUserAPI struct {
	ssov1.UnimplementedUserServer
	User *UserImpl
}

func Register(gRPC *grpc.Server, user *UserImpl) {
	ssov1.RegisterUserServer(
		gRPC,
		&ServerUserAPI{User: user},
	)
}

func (s *ServerUserAPI) GetUser(ctx context.Context, req *ssov1.GetUserRequest) (*ssov1.GetUserResponse, error) {
	resp, err := s.User.GetUserData(ctx, req)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	return resp, nil
}

func (s *ServerUserAPI) ChangeProfilePhoto(ctx context.Context, req *ssov1.ChangeProfilePhotoRequest) (*ssov1.ChangeProfilePhotoResponse, error) {
	resp, err := s.User.ChangeProfilePhoto(ctx, req)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	return resp, nil
}

func (s *ServerUserAPI) DeleteUser(ctx context.Context, req *ssov1.UserDeleteRequest) (*ssov1.UserDeleteResponse, error) {
	deleteStatus, err := s.User.DeleteUser(ctx, req)
	if err != nil {

	}

	return &ssov1.UserDeleteResponse{Status: deleteStatus}, nil
}

func (s *ServerUserAPI) GetUsersData(ctx context.Context, req *ssov1.UserName) (*ssov1.GetUsersDataResponse, error) {
	deleteStatus, err := s.User.GetUsersData(ctx, req)
	if err != nil {

	}

	return deleteStatus, nil
}

func (s *ServerUserAPI) ChangePassword(ctx context.Context, req *ssov1.NewPassword) (*ssov1.Empty, error) {
	err := s.User.ChangePassword(ctx, req)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *ServerUserAPI) ChangeEmail(ctx context.Context, req *ssov1.Email) (*ssov1.GetUserResponse, error) {
	err := s.User.ChangeEmail(ctx, req)
	if err != nil {
		return nil, err
	}

	resp, err := s.User.GetUserData(ctx, &ssov1.GetUserRequest{UserName: req.UserName})
	if err != nil {
		return nil, err
	}

	return resp, nil
}
