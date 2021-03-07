package grpc_load_balancer_application

import (
	"github.com/Riphal/grpc-load-balancer-application/common/storage/postgres"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/server"
)

type App struct {
	Server	*server.Server

	DB 		*postgres.DB
}

func NewApp() *App {
	// TO-DO: Implement
	return &App{}
}
