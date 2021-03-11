package main

import (
	core "github.com/Riphal/grpc-load-balancer-application"
	loadBalancerProto "github.com/Riphal/grpc-load-balancer-application/common/proto/loadBalancer"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/controller"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/service"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/service/loadBalancer"
	loadBalancerStorage "github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/storage/redis/loadBalancer"
	"google.golang.org/grpc"
)

func registerGRPCServer(server *grpc.Server, app *core.App) {
	// Init global config
	serviceConfig := &service.Config{}
	controllerConfig := &controller.Config{}

	// Init Services
	var loadBalancerService loadBalancer.Service = loadBalancer.NewServiceImplementation(&loadBalancer.Config{
		Config: 				serviceConfig,
		LoadBalancerStorage:	loadBalancerStorage.NewRedisImplementation(app.Redis),
	})

	// Init controllers
	loadBalancerController := controller.NewLoadBalancerController(controllerConfig, loadBalancerService)

	loadBalancerProto.RegisterLoadBalancerServiceServer(server, loadBalancerController)
}
