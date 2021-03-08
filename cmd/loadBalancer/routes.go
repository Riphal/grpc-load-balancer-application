package main

import (
	core "github.com/Riphal/grpc-load-balancer-application"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/service"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/service/auth"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/service/auth/jwt"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/storage/postgres/account"
	authStorage "github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/storage/redis/auth"
)

func registerRoutes(app *core.App) {
	var authService auth.Service = auth.NewServiceImplementation(&auth.Config{
		Config: 		&service.Config{},
		AuthStorage: 	authStorage.NewRedisImplementation(app.Redis),
		AccountStorage:	account.NewPGStorageImplementation(app.DB),
		JwtService: 	jwt.NewServiceImplementation(),
	})

	_ = authService
}
