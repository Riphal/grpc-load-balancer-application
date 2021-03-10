package storage

import (
	"context"
	"github.com/Riphal/grpc-load-balancer-application/common/errors"
	"github.com/Riphal/grpc-load-balancer-application/common/model/bankAccount"
	"github.com/Riphal/grpc-load-balancer-application/common/model/expense"
)

type BankAccount interface {
	GetBankAccounts(ctx context.Context, accountID string) ([]*bankAccount.BankAccount, errors.Error)
	GetBankAccount(ctx context.Context, id string) (*bankAccount.BankAccountBalance, errors.Error)
	CreateBankAccount(ctx context.Context, bankAccount *bankAccount.BankAccount) errors.Error
	DeleteBankAccount(ctx context.Context, id string) errors.Error
}

type Expense interface {
	GetExpenses(ctx context.Context, bankAccountID string) ([]*expense.Expense, errors.Error)
	CreateExpense(ctx context.Context, expense *expense.Expense) errors.Error
	DeleteExpense(ctx context.Context, id string) errors.Error
}
