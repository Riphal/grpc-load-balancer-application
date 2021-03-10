package storage

import (
	"context"
	"github.com/Riphal/grpc-load-balancer-application/common/errors"
	"github.com/Riphal/grpc-load-balancer-application/common/model/bankAccount"
)

type BankAccount interface {
	//GetBankAccounts(ctx context.Context) (*bankAccount.BankAccount, errors.Error)
	GetBankAccount(ctx context.Context, id string) (*bankAccount.BankAccount, errors.Error)
	CreateBankAccount(ctx context.Context, bankAccount *bankAccount.BankAccount) errors.Error
	//DeleteBankAccount(ctx context.Context, id string) errors.Error
}
