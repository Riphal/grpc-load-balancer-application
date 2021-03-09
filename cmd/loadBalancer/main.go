package main

import (
	"fmt"
	"log"

	core "github.com/Riphal/grpc-load-balancer-application"
	"github.com/Riphal/grpc-load-balancer-application/common/config"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/server"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/server/router"
	pkgerrors "github.com/pkg/errors"
)

func main() {
	app := core.NewApp()

	s := mustInitServer()

	registerRoutes(s, app)

	if err := s.ListenAndServe(); err != nil {
		log.Fatalln(pkgerrors.Wrap(err, "failed to run server"))
	}
	//conn, err := loadBalancer.Dial(":9001", loadBalancer.WithInsecure())
	//if err != nil {
	//	log.Fatalf("Could't connect: %s", err)
	//}
	//
	//c := chat.NewChatServiceClient(conn)
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
