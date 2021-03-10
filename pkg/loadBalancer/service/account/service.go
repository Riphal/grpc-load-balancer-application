package account

import (
	"context"

	"github.com/Riphal/grpc-load-balancer-application/common/errors"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/model/response"
)

type Service interface {
	GetAccount(ctx context.Context, id string) (*response.AccountResponse, errors.Error)
}
