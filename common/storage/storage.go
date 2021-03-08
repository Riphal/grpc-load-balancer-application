package storage

import (
	"context"

	"github.com/Riphal/grpc-load-balancer-application/common/errors"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/model/account"
)

type Account interface {
	GetAccount(ctx context.Context, id string) (*account.Account, errors.Error)
	CreateAccount(ctx context.Context, account *account.Account) errors.Error
	DeleteAccount(ctx context.Context, id string) errors.Error
}
