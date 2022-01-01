package main

import (
	"context"
	"github.com/erikjiang/grpc_demo/helper"
	"github.com/erikjiang/grpc_demo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type UserService struct {
	proto.UnimplementedUserServiceServer
}

func (us *UserService) GetUserInfo(ctx context.Context, in *proto.UserRequest) (*proto.UserResponse, error) {
	log.Printf("user id: %d\n", in.UserId)
	return &proto.UserResponse{
		UserId:   in.UserId,
		UserName: "Jason",
		UserDesc: "Jason is a engineer",
	}, nil
}

func main() {
	// 服务端双向认证设置:
	creds := helper.GetServerCreds()

	// 1. create grpc server (set credentials)
	server := grpc.NewServer(grpc.Creds(creds))
	// 1.1. support debug by grpcui
	reflection.Register(server)
	// 2. register grpc user service
	proto.RegisterUserServiceServer(server, &UserService{})
	// 3. run tcp server and listen address
	lis, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		log.Fatalf("listen failed, err: %s\n", err.Error())
	}
	log.Printf("listen address: %s\n", lis.Addr())
	// 4. grpc server accepts tcp listener
	if err := server.Serve(lis); err != nil {
		log.Fatalf("serve failed, err: %s", err.Error())
	}

}
