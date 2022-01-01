package main

import (
	"context"
	"github.com/erikjiang/grpc_demo/helper"
	"github.com/erikjiang/grpc_demo/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	// 客户端双向认证设置:
	creds := helper.GetClientCreds()

	//conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure(), grpc.WithBlock())
	conn, err := grpc.Dial("localhost:8081", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("dial failed, err: %s\n", err.Error())
	}
	client := proto.NewUserServiceClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	resp, err := client.GetUserInfo(ctx, &proto.UserRequest{
		UserId: 987654,
	})
	if err != nil {
		log.Fatalf("get user info failed, err: %s\n", err.Error())
	}
	log.Printf("get user info, response: %v\n", resp)

}
