package main

import (
	"fmt"
	"log"
	"net"

	core "github.com/Riphal/grpc-load-balancer-application"
	"github.com/Riphal/grpc-load-balancer-application/common/config"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/server"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/server/router"
	pkgerrors "github.com/pkg/errors"
	"google.golang.org/grpc"
)

func main() {
	app := core.NewApp()

	s := mustInitServer()
	listener, grpcServer := mustInitGRPC()

	registerGRPCServer(grpcServer, app)
	registerRoutes(s, app)

	go listenAndServeConnectionGRPCErrors(listener, grpcServer)
	listenAndServeConnectionRouterErrors(s)
}

func mustInitServer() *server.Server {
	r := router.New()
	port := ":" + config.GetEnv("PORT", "3001")

	serverConfig := &server.Config{
		Address: port,
		Router:  r,
	}

	log.Println(fmt.Sprintf("🚀 server listen on port %v", port))

	return server.New(serverConfig)
}

func listenAndServeConnectionRouterErrors(s *server.Server) {
	if err := s.ListenAndServe(); err != nil {
		log.Fatalln(pkgerrors.Wrap(err, "failed to run server"))
	}
}

func mustInitGRPC () (net.Listener, *grpc.Server) {
	port := "4090"

	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalf("faild to listen on port %v: %v", port, err)
	}

	grpcServer := grpc.NewServer()

	log.Println(fmt.Sprintf("🚀 loadBalancer gRPC server listen on %v", listener.Addr()))

	return listener, grpcServer
}

func listenAndServeConnectionGRPCErrors(listener net.Listener, grpcServer *grpc.Server) {
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("faild to serve gRPC server over %v: %v", listener.Addr(), err)
	}
}
