package main

import (
	core "github.com/Riphal/grpc-load-balancer-application"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/controller"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/middleware"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/service"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/service/auth"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/service/auth/jwt"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/storage/postgres/account"
	authStorage "github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/storage/redis/auth"
)

func registerRoutes(app *core.App) {
	// Init services
	var authService auth.Service = auth.NewServiceImplementation(&auth.Config{
		Config: 		&service.Config{},
		AuthStorage: 	authStorage.NewRedisImplementation(app.Redis),
		AccountStorage:	account.NewPGStorageImplementation(app.DB),
		JwtService: 	jwt.NewServiceImplementation(),
	})

	// Init controllers
	authController := controller.NewAuthController(&controller.Config{}, authService)

	// Init router
	api := app.Server.Router.Group("/api/v1")

	// Auth routes
	api.POST("/register", authController.Register)
	api.POST("/login", authController.Login)
	api.POST("/logout", middleware.AuthorizeJWT(authService), authController.Logout)
}
