package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"github.com/erikjiang/grpc_demo/certs"
	"github.com/erikjiang/grpc_demo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
)

func main() {
	// 客户端双向认证设置:
	// load client key & pem
	cert, err := tls.LoadX509KeyPair(certs.Path("client/client.pem"), certs.Path("client/client.key"))
	if err != nil {
		log.Fatalf("load x509 cert file failed, err: %s\n", err.Error())
	}
	// create cert pool and  append ca pem
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(certs.Path("ca/ca.pem"))
	if err != nil {
		log.Fatalf("read ca file failed, err: %s\n", err.Error())
	}
	certPool.AppendCertsFromPEM(ca)
	// set credentials
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   "localhost", // 与证书中的"CN"即:Common Name相同
		RootCAs:      certPool,
	})

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
