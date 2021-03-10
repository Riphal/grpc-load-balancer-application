package bankAccount

import (
	"context"
	"github.com/Riphal/grpc-load-balancer-application/common/model/bankAccount"
	bankAccountProto "github.com/Riphal/grpc-load-balancer-application/common/proto/bankAccount"
	"github.com/Riphal/grpc-load-balancer-application/pkg/worker/service"
	"github.com/Riphal/grpc-load-balancer-application/pkg/worker/storage"
)

type Config struct {
	*service.Config
	BankAccountStorage	storage.BankAccount
}

type ServiceImplementation struct {
	*service.Service
	bankAccountStorage	storage.BankAccount
}

func NewServiceImplementation(config *Config) *ServiceImplementation {
	return &ServiceImplementation{
		Service:			service.New(config.Config),
		bankAccountStorage:	config.BankAccountStorage,
	}
}

func (si *ServiceImplementation) CreateBankAccount(ctx context.Context, bankAccount *bankAccount.BankAccount) *bankAccountProto.Error {
	err := si.bankAccountStorage.CreateBankAccount(ctx, bankAccount)
	if err.IsNotNil() {
		return bankAccountProto.ErrorToProto(&err)
	}

	return bankAccountProto.NilErrorToProto()
}
