package main

import (
	"context"
	"github.com/erikjiang/grpc_demo/helper"
	"github.com/erikjiang/grpc_demo/proto"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(helper.GetClientCreds())}
	err := proto.RegisterUserServiceHandlerFromEndpoint(ctx, mux, "localhost:8081", opts)
	if err != nil {
		log.Fatalf("register user service handler failed, err: %s\n", err.Error())
	}
	err = proto.RegisterTeamServiceHandlerFromEndpoint(ctx, mux, "localhost:8081", opts)
	if err != nil {
		log.Fatalf("register team service handler failed, err: %s\n", err.Error())
	}

	http.ListenAndServe("localhost:8080", mux)
}
