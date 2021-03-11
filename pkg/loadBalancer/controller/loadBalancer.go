package controller

import (
	"context"

	loadBalancerProto "github.com/Riphal/grpc-load-balancer-application/common/proto/loadBalancer"
	loadBalancerService "github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/service/loadBalancer"
	"google.golang.org/protobuf/types/known/emptypb"
)

type LoadBalancerController struct {
	*Controller
	loadBalancerService loadBalancerService.Service
}

func NewLoadBalancerController(config *Config, loadBalancerService loadBalancerService.Service) *LoadBalancerController {
	return &LoadBalancerController{
		Controller:		NewController(config),
		loadBalancerService:	loadBalancerService,
	}
}

func (lbc LoadBalancerController) Register(ctx context.Context, request *loadBalancerProto.Request) (*emptypb.Empty, error) {
	_ = lbc.loadBalancerService.RegisterWorker(ctx, request.Addr)

	return &emptypb.Empty{}, nil
}

func (lbc LoadBalancerController) DeRegister(ctx context.Context, request *loadBalancerProto.Request) (*emptypb.Empty, error) {
	_ = lbc.loadBalancerService.DeRegisterWorker(ctx, request.Addr)

	return &emptypb.Empty{}, nil
}
