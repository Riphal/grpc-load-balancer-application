package auth

import (
	"context"

	"github.com/Riphal/grpc-load-balancer-application/common/errors"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/model/account"
	jwtModel "github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/model/jwt"
)

type Service interface {
	Register(ctx context.Context, account *account.Account) (string, errors.Error)
	Login(ctx context.Context, email, password string) (string, errors.Error)
	Logout(ctx context.Context, token string) errors.Error
	ValidateToken(ctx context.Context, token string) (*jwtModel.Claims, errors.Error)
}
