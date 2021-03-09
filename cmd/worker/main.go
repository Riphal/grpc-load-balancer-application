package main

import (
	"fmt"
	"log"
	"net"
	"os"

	core "github.com/Riphal/grpc-load-balancer-application"
	"google.golang.org/grpc"
)

func main() {
	app := core.NewApp()
	listener, grpcServer := mustInitGRPC()

	// register services to gRPC server
	_ = app

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("faild to serve gRPC server over %v: %v", listener.Addr(), err)
	}
	log.Println(fmt.Sprintf("🚀 gRPC listen on %v", listener.Addr()))
}

func mustInitGRPC () (net.Listener, *grpc.Server) {
	port := os.Getenv("PORT")

	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalf("faild to listen on port %v: %v", port, err)
	}

	grpcServer := grpc.NewServer()

	return listener, grpcServer
}
