package storage

import (
	"context"

	"github.com/Riphal/grpc-load-balancer-application/common/errors"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/model/account"
)

type Auth interface {
	IsBlacklisted(ctx context.Context, token string) (bool, errors.Error)
	SetBlacklistToken(ctx context.Context, token string) errors.Error
}

type Account interface {
	GetAccount(ctx context.Context, email string) (*account.Account, errors.Error)
	CreateAccount(ctx context.Context, account *account.Account) errors.Error
	DeleteAccount(ctx context.Context, id string) errors.Error
}
