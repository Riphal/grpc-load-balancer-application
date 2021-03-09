package grpc

import (
	"context"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/model/expense"

	"github.com/Riphal/grpc-load-balancer-application/common/errors"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/model/bankAccount"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/model/response"
)

type Service interface {
	// Account service
	GetAccount(ctx context.Context, id string) (*response.AccountResponse, errors.Error)

	// Bank accounts service
	GetBankAccounts(ctx context.Context) (*response.BankAccountsResponse, errors.Error)
	GetBankAccount(ctx context.Context, id string) (*response.BankAccountResponse, errors.Error)
	CreateBankAccount(ctx context.Context, bankAccount *bankAccount.BankAccount) errors.Error
	DeleteBankAccount(ctx context.Context, id string) errors.Error

	// Bank account expenses
	GetExpenses(ctx context.Context, bankAccountID string) (*response.ExpensesResponse, errors.Error)
	CreateExpense(ctx context.Context, expense *expense.Expense) errors.Error
	DeleteExpense(ctx context.Context, id string) errors.Error
}
