package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"github.com/erikjiang/grpc_demo/certs"
	"github.com/erikjiang/grpc_demo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"io/ioutil"
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
	// load server key & pem
	cert, err := tls.LoadX509KeyPair(certs.Path("server/server.pem"), certs.Path("server/server.key"))
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
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	})

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
