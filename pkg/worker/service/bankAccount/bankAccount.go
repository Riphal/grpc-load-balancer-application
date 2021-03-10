package bankAccount

import (
	"context"
	"github.com/Riphal/grpc-load-balancer-application/common/errors"
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


func (si *ServiceImplementation) GetBankAccounts(ctx context.Context, accountID string) *bankAccountProto.BankAccountsReply {
	bankAccounts, err := si.bankAccountStorage.GetBankAccounts(ctx, accountID)
	if err.IsNotNil() {
		return &bankAccountProto.BankAccountsReply{
			Error: bankAccountProto.ErrorToProto(err),
		}
	}

	return bankAccountProto.GetBankAccountsReplyToProto(bankAccounts, err)
}

func (si *ServiceImplementation) GetBankAccount(ctx context.Context, id string) *bankAccountProto.BankAccountReply {
	bankAcc, err := si.bankAccountStorage.GetBankAccount(ctx, id)
	if err.IsNotNil() {
		return &bankAccountProto.BankAccountReply{
			Error: bankAccountProto.ErrorToProto(err),
		}
	}

	return bankAccountProto.GetBankAccountReplyToProto(bankAcc, err)
}

func (si *ServiceImplementation) CreateBankAccount(ctx context.Context, bankAccount *bankAccount.BankAccount) *bankAccountProto.Error {
	err := si.bankAccountStorage.CreateBankAccount(ctx, bankAccount)
	if err.IsNotNil() {
		return bankAccountProto.ErrorToProto(err)
	}

	return bankAccountProto.ErrorToProto(errors.Nil())
}

func (si *ServiceImplementation) DeleteBankAccount(ctx context.Context, id string) *bankAccountProto.Error {
	err := si.bankAccountStorage.DeleteBankAccount(ctx, id)
	if err.IsNotNil() {
		return bankAccountProto.ErrorToProto(err)
	}

	return bankAccountProto.ErrorToProto(errors.Nil())
}
