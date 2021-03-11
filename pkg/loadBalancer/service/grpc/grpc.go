package grpc

import (
	"context"
	"github.com/Riphal/grpc-load-balancer-application/common/errors"
	"github.com/Riphal/grpc-load-balancer-application/common/model/bankAccount"
	"github.com/Riphal/grpc-load-balancer-application/common/model/expense"
	bankAccountProto "github.com/Riphal/grpc-load-balancer-application/common/proto/bankAccount"
	expenseProto "github.com/Riphal/grpc-load-balancer-application/common/proto/expense"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/model/response"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/service"
	loadBalancerService "github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/service/loadBalancer"
)

type Config struct {
	*service.Config
	LoadBalancerService loadBalancerService.Service
}

type ServiceImplementation struct {
	*service.Service
	loadBalancerService loadBalancerService.Service
}

func NewServiceImplementation(config *Config) *ServiceImplementation {
	return &ServiceImplementation{
		Service:        		service.New(config.Config),
		loadBalancerService:	config.LoadBalancerService,
	}
}

func (si *ServiceImplementation) GetBankAccounts(ctx context.Context, accountID string) ([]response.BankAccountResponse, errors.Error) {
	bankAccountsReq := bankAccountProto.GetBankAccountsRequestToProto(accountID)

	grpcClient, err := si.loadBalancerService.GRPCClient(ctx)
	if err.IsNotNil() {
		return nil, err
	}

	protoReply, _ := grpcClient.BankAccountServiceClient.GetBankAccounts(ctx, bankAccountsReq)
	err = bankAccountProto.ErrorToModel(protoReply.Error)
	if err.IsNotNil() {
		return nil, err
	}

	return bankAccountProto.GetBankAccountsReplyToModel(protoReply)
}

func (si *ServiceImplementation) GetBankAccount(ctx context.Context, id string) (*bankAccount.BankAccountBalance, errors.Error) {
	bankAccReq := bankAccountProto.GetBankAccountRequestToProto(id)

	grpcClient, err := si.loadBalancerService.GRPCClient(ctx)
	if err.IsNotNil() {
		return nil, err
	}

	protoReply, _ := grpcClient.BankAccountServiceClient.GetBankAccount(ctx, bankAccReq)
	err = bankAccountProto.ErrorToModel(protoReply.Error)
	if err.IsNotNil() {
		return nil, err
	}

	return bankAccountProto.GetBankAccountReplyToModel(protoReply)
}

func (si *ServiceImplementation) CreateBankAccount(ctx context.Context, bankAccount *bankAccount.BankAccount) errors.Error {
	bankAccReq := bankAccountProto.CreateBankAccountRequestToProto(bankAccount)

	grpcClient, err := si.loadBalancerService.GRPCClient(ctx)
	if err.IsNotNil() {
		return err
	}

	protoErr, _ := grpcClient.BankAccountServiceClient.CreateBankAccount(ctx, bankAccReq)
	err = bankAccountProto.ErrorToModel(protoErr)
	if err.IsNotNil() {
		return err
	}

	return errors.Nil()
}

func (si *ServiceImplementation) DeleteBankAccount(ctx context.Context, id string) errors.Error {
	bankAccReq := bankAccountProto.GetBankAccountRequestToProto(id)

	grpcClient, err := si.loadBalancerService.GRPCClient(ctx)
	if err.IsNotNil() {
		return err
	}

	protoErr, _ := grpcClient.BankAccountServiceClient.DeleteBankAccount(ctx, bankAccReq)
	err = bankAccountProto.ErrorToModel(protoErr)
	if err.IsNotNil() {
		return err
	}

	return errors.Nil()
}


func (si *ServiceImplementation) GetExpenses(ctx context.Context, bankAccountID string) ([]expense.Expense, errors.Error) {
	expensesReq := expenseProto.GetExpensesRequestToProto(bankAccountID)

	grpcClient, err := si.loadBalancerService.GRPCClient(ctx)
	if err.IsNotNil() {
		return nil, err
	}

	protoReply, _ := grpcClient.ExpenseServiceClient.GetExpenses(ctx, expensesReq)
	err = expenseProto.ErrorToModel(protoReply.Error)
	if err.IsNotNil() {
		return nil, err
	}

	return expenseProto.GetExpensesReplyToModel(protoReply)
}

func (si *ServiceImplementation) GetExpense(ctx context.Context, id string) (*expense.Expense, errors.Error) {
	expenseReq := expenseProto.GetExpenseRequestToProto(id)

	grpcClient, err := si.loadBalancerService.GRPCClient(ctx)
	if err.IsNotNil() {
		return nil, err
	}

	protoReply, _ := grpcClient.ExpenseServiceClient.GetExpense(ctx, expenseReq)
	err = expenseProto.ErrorToModel(protoReply.Error)
	if err.IsNotNil() {
		return nil, err
	}

	return expenseProto.GetExpenseReplyToModel(protoReply)
}

func (si *ServiceImplementation) CreateExpense(ctx context.Context, expense *expense.Expense) errors.Error {
	expenseReq := expenseProto.CreateExpenseRequestToProto(expense)

	grpcClient, err := si.loadBalancerService.GRPCClient(ctx)
	if err.IsNotNil() {
		return err
	}

	protoErr, _ := grpcClient.ExpenseServiceClient.CreateExpense(ctx, expenseReq)
	err = expenseProto.ErrorToModel(protoErr)
	if err.IsNotNil() {
		return err
	}

	return errors.Nil()
}

func (si *ServiceImplementation) DeleteExpense(ctx context.Context, id string) errors.Error {
	expenseReq := expenseProto.DeleteExpenseRequestToProto(id)

	grpcClient, err := si.loadBalancerService.GRPCClient(ctx)
	if err.IsNotNil() {
		return err
	}

	protoErr, _ := grpcClient.ExpenseServiceClient.DeleteExpense(ctx, expenseReq)
	err = expenseProto.ErrorToModel(protoErr)
	if err.IsNotNil() {
		return err
	}

	return errors.Nil()
}
