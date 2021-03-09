package grpc

import (
	"context"
	"github.com/Riphal/grpc-load-balancer-application/common/errors"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/model/bankAccount"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/model/expense"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/model/response"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/service"
	lbstorage "github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/storage"
)

type Config struct {
	*service.Config
	LoadBalancerStorage	lbstorage.LoadBalancer
}

type ServiceImplementation struct {
	*service.Service
	LoadBalancerStorage	lbstorage.LoadBalancer
}

func NewServiceImplementation(config *Config) *ServiceImplementation {
	return &ServiceImplementation{
		Service:        		service.New(config.Config),
		LoadBalancerStorage:	config.LoadBalancerStorage,
	}
}


func (si *ServiceImplementation) GetAccount(ctx context.Context, id string) (*response.AccountResponse, errors.Error) {
	// TO-DO: Implement
	return &response.AccountResponse{}, errors.Nil()
}


func (si *ServiceImplementation) GetBankAccounts(ctx context.Context) (*response.BankAccountsResponse, errors.Error) {
	// TO-DO: Implement
	return &response.BankAccountsResponse{}, errors.Nil()
}

func (si *ServiceImplementation) GetBankAccount(ctx context.Context, id string) (*response.BankAccountResponse, errors.Error) {
	// TO-DO: Implement
	return &response.BankAccountResponse{}, errors.Nil()
}

func (si *ServiceImplementation) CreateBankAccount(ctx context.Context, bankAccount *bankAccount.BankAccount) errors.Error {
	// TO-DO: Implement
	return errors.Nil()
}

func (si *ServiceImplementation) DeleteBankAccount(ctx context.Context, id string) errors.Error {
	// TO-DO: Implement
	return errors.Nil()
}


func (si *ServiceImplementation) GetExpenses(ctx context.Context, bankAccountID string) (*response.ExpensesResponse, errors.Error) {
	// TO-DO: Implement
	return &response.ExpensesResponse{}, errors.Nil()
}

func (si *ServiceImplementation) CreateExpense(ctx context.Context, expense *expense.Expense) errors.Error {
	// TO-DO: Implement
	return errors.Nil()
}

func (si *ServiceImplementation) DeleteExpense(ctx context.Context, id string) errors.Error {
	// TO-DO: Implement
	return errors.Nil()
}
