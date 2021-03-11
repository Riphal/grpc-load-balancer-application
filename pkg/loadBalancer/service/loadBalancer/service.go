package loadBalancer

import (
	"context"
	"github.com/Riphal/grpc-load-balancer-application/common/errors"
)

type Service interface {
	RegisterWorker(ctx context.Context, addr string) errors.Error
	DeRegisterWorker(ctx context.Context, addr string) errors.Error
	GRPCClient(ctx context.Context) (*GRPCClient, errors.Error)
}
