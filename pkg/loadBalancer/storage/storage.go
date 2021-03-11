package storage

import (
	"context"

	"github.com/Riphal/grpc-load-balancer-application/common/errors"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/model/account"
)

type Account interface {
	GetAccount(ctx context.Context, id string) (*account.Account, errors.Error)
	GetAccountWithEmail(ctx context.Context, email string) (*account.Account, errors.Error)
	CreateAccount(ctx context.Context, account *account.Account) errors.Error
}

type Auth interface {
	IsBlacklisted(ctx context.Context, token string) (bool, errors.Error)
	SetBlacklistToken(ctx context.Context, token string) errors.Error
}

type LoadBalancer interface {
	IncrCounter(ctx context.Context) (int64, errors.Error)
	GetWorkers(ctx context.Context) ([]string, errors.Error)
	RegisterWorker(ctx context.Context, addr string) errors.Error
	DeRegisterWorker(ctx context.Context, addr string) errors.Error
}
