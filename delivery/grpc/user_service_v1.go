package grpc

import (
	"context"

	accountV1 "github.com/itmrchow/microservice-proto/account/v1"
)

type UserServiceV1 struct {
	accountV1.UnimplementedUserServiceServer
}

func NewUserServiceV1() *UserServiceV1 {
	return &UserServiceV1{}
}

func (s *UserServiceV1) GetUser(ctx context.Context, req *accountV1.GetUserRequest) (*accountV1.GetUserResponse, error) {

	resp := &accountV1.GetUserResponse{
		Id:       req.Id,
		Username: "test",
	}

	return resp, nil
}
