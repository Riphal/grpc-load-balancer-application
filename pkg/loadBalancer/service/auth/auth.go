package auth

import (
	"context"

	"github.com/Riphal/grpc-load-balancer-application/common/errors"
	"github.com/Riphal/grpc-load-balancer-application/common/model/account"
	"github.com/Riphal/grpc-load-balancer-application/common/storage"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/service"
	lbstorage "github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/storage"
)

type Config struct {
	*service.Config
	AuthStorage lbstorage.Auth
	AccountStorage storage.Account
}

type ServiceImplementation struct {
	*service.Service
	authStorage lbstorage.Auth
	accountStorage storage.Account
}

func NewServiceImplementation(config *Config) *ServiceImplementation {
	return &ServiceImplementation{
		Service:        service.New(config.Config),
		authStorage: config.AuthStorage,
		accountStorage: config.AccountStorage,
	}
}

func (si *ServiceImplementation) Register(ctx context.Context, account *account.Account) (string, errors.Error) {
	// TO-DO: Implement
	return "", errors.Nil()
}

func (si *ServiceImplementation) Login(ctx context.Context, account *account.Account) (string, errors.Error) {
	// TO-DO: Implement
	return "", errors.Nil()
}

func (si *ServiceImplementation) Logout(ctx context.Context, token string) errors.Error {
	// TO-DO: Implement
	return errors.Nil()
}
