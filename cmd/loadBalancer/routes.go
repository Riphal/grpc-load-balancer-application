package main

import (
	core "github.com/Riphal/grpc-load-balancer-application"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/controller"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/middleware"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/server"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/service"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/service/auth"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/service/auth/jwt"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/service/grpc"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/storage/postgres/account"
	authStorage "github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/storage/redis/auth"
	loadBalancerStorage "github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/storage/redis/loadBalancer"
)

func registerRoutes(server *server.Server, app *core.App) {
	// Init global config
	serviceConfig := &service.Config{}
	controllerConfig := &controller.Config{}

	// Init services
	var authService auth.Service = auth.NewServiceImplementation(&auth.Config{
		Config: 		serviceConfig,
		AuthStorage: 	authStorage.NewRedisImplementation(app.Redis),
		AccountStorage:	account.NewPGStorageImplementation(app.DB),
		JwtService: 	jwt.NewServiceImplementation(),
	})

	var grpcService grpc.Service = grpc.NewServiceImplementation(&grpc.Config{
		Config: 				serviceConfig,
		LoadBalancerStorage: 	loadBalancerStorage.NewRedisImplementation(app.Redis),
	})


	// Init controllers
	authController := controller.NewAuthController(controllerConfig, authService)
	grpcController := controller.NewGRPCController(controllerConfig, grpcService)


	// Init router
	api := server.Router.Group("/api/v1")

	// Auth routes
	api.POST("/register", authController.Register)
	api.POST("/login", authController.Login)
	api.POST("/logout", middleware.AuthorizeJWT(authService), authController.Logout)


	// Account routes
	accountRouter := api.Group("/account", middleware.AuthorizeJWT(authService))

	accountRouter.GET("/", grpcController.GetAccount)


	// Bank accounts routes
	bankAccountsRouter := accountRouter.Group("/bank-accounts")

	bankAccountsRouter.GET("/", grpcController.GetBankAccounts)
	bankAccountsRouter.GET("/:bank_account_id", grpcController.GetBankAccount)
	bankAccountsRouter.POST("/", grpcController.CreateBankAccount)
	bankAccountsRouter.DELETE("/:bank_account_id", grpcController.DeleteBankAccount)


	// Bank expenses routes
	expensesRouter := bankAccountsRouter.Group("/expenses")

	expensesRouter.GET("/", grpcController.GetExpenses)
	expensesRouter.POST("/", grpcController.CreateExpense)
	expensesRouter.DELETE("/:expense_id", grpcController.DeleteExpense)
}
