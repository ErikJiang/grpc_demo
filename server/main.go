package main

import (
	"github.com/erikjiang/grpc_demo/helper"
	"github.com/erikjiang/grpc_demo/proto"
	"github.com/erikjiang/grpc_demo/server/svc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	// 服务端双向认证设置:
	creds := helper.GetServerCreds()

	// 1. create grpc server (set credentials)
	s := grpc.NewServer(grpc.Creds(creds))
	// 1.1. support debug by grpcui
	reflection.Register(s)
	// 2. register grpc user&team service
	proto.RegisterUserServiceServer(s, &svc.UserService{})
	proto.RegisterTeamServiceServer(s, &svc.TeamService{})
	// 3. run tcp server and listen address
	lis, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		log.Fatalf("listen failed, err: %s\n", err.Error())
	}
	log.Printf("listen address: %s\n", lis.Addr())
	// 4. grpc server accepts tcp listener
	if err := s.Serve(lis); err != nil {
		log.Fatalf("serve failed, err: %s", err.Error())
	}

}
