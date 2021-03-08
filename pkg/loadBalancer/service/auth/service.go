package auth

import (
	"context"

	"github.com/Riphal/grpc-load-balancer-application/common/errors"
	"github.com/Riphal/grpc-load-balancer-application/common/model/account"
)

type Service interface {
	Register(ctx context.Context, account *account.Account) (string, errors.Error)
	Login(ctx context.Context, account *account.Account) (string, errors.Error)
	Logout(ctx context.Context, token string) errors.Error
}
