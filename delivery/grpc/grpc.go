package grpc

import (
	"net"

	accountV1 "github.com/itmrchow/microservice-proto/account/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func RunGrpcServer(port string) error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	// middleware
	opts := []grpc.ServerOption{}

	s := grpc.NewServer(opts...)

	accountV1.RegisterUserServiceServer(s, NewUserServiceV1())

	reflection.Register(s)
	err = s.Serve(lis)

	return err
}
