package account

import (
	"context"

	"github.com/Riphal/grpc-load-balancer-application/common/errors"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/model/response"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/service"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/storage"
)

type Config struct {
	*service.Config
	AccountStorage 	storage.Account
}

type ServiceImplementation struct {
	*service.Service
	accountStorage 	storage.Account
}

func NewServiceImplementation(config *Config) *ServiceImplementation {
	return &ServiceImplementation{
		Service:        service.New(config.Config),
		accountStorage: config.AccountStorage,
	}
}

func (si *ServiceImplementation) GetAccount(ctx context.Context, id string) (*response.AccountResponse, errors.Error) {
	acc, err := si.accountStorage.GetAccount(ctx, id)
	if err.IsNotNil() {
		return nil, err
	}

	return &response.AccountResponse{
		Email: acc.Email,
		FirstName: acc.FirstName,
		LastName: acc.LastName,
	}, errors.Nil()
}
