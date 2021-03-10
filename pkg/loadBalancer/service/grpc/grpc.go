package grpc

import (
	"context"
	"log"

	"github.com/Riphal/grpc-load-balancer-application/common/errors"
	"github.com/Riphal/grpc-load-balancer-application/common/model/bankAccount"
	"github.com/Riphal/grpc-load-balancer-application/common/model/expense"
	bankAccountProto "github.com/Riphal/grpc-load-balancer-application/common/proto/bankAccount"
	expenseProto "github.com/Riphal/grpc-load-balancer-application/common/proto/expense"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/model/response"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/service"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/storage"
	"google.golang.org/grpc"
)

type Config struct {
	*service.Config
	LoadBalancerStorage	storage.LoadBalancer
}

type GRPCClient struct {
	BankAccountServiceClient 	bankAccountProto.BankAccountServiceClient
	ExpenseServiceClient		expenseProto.ExpenseServiceClient
}

type ServiceImplementation struct {
	*service.Service
	gRPCClient 				*GRPCClient
	loadBalancerStorage		storage.LoadBalancer
}

func NewServiceImplementation(config *Config) *ServiceImplementation {
	return &ServiceImplementation{
		Service:        		service.New(config.Config),
		gRPCClient:				mustInitGRPCClientConn(),
		loadBalancerStorage:	config.LoadBalancerStorage,
	}
}


// TO-DO: Move to load balancer service
func mustInitGRPCClientConn() *GRPCClient {
	conn, err := grpc.Dial(
		":9001",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("could't connect: %s", err)
	}

	return &GRPCClient{
		BankAccountServiceClient: 	bankAccountProto.NewBankAccountServiceClient(conn),
		ExpenseServiceClient:		expenseProto.NewExpenseServiceClient(conn),
	}
}

func (si *ServiceImplementation) GetBankAccounts(ctx context.Context, accountID string) ([]response.BankAccountResponse, errors.Error) {
	bankAccountsReq := bankAccountProto.GetBankAccountsRequestToProto(accountID)

	protoReply, _ := si.gRPCClient.BankAccountServiceClient.GetBankAccounts(ctx, bankAccountsReq)
	err := bankAccountProto.ErrorToModel(protoReply.Error)
	if err.IsNotNil() {
		return nil, err
	}

	return bankAccountProto.GetBankAccountsReplyToModel(protoReply)
}

func (si *ServiceImplementation) GetBankAccount(ctx context.Context, id string) (*response.BankAccountResponse, errors.Error) {
	bankAccReq := bankAccountProto.GetBankAccountRequestToProto(id)

	protoReply, _ := si.gRPCClient.BankAccountServiceClient.GetBankAccount(ctx, bankAccReq)
	err := bankAccountProto.ErrorToModel(protoReply.Error)
	if err.IsNotNil() {
		return nil, err
	}

	return bankAccountProto.GetBankAccountReplyToModel(protoReply)
}

func (si *ServiceImplementation) CreateBankAccount(ctx context.Context, bankAccount *bankAccount.BankAccount) errors.Error {
	bankAccReq := bankAccountProto.CreateBankAccountRequestToProto(bankAccount)

	protoErr, _ := si.gRPCClient.BankAccountServiceClient.CreateBankAccount(ctx, bankAccReq)
	err := bankAccountProto.ErrorToModel(protoErr)
	if err.IsNotNil() {
		return err
	}

	return errors.Nil()
}

func (si *ServiceImplementation) DeleteBankAccount(ctx context.Context, id string) errors.Error {
	bankAccReq := bankAccountProto.GetBankAccountRequestToProto(id)

	protoErr, _ := si.gRPCClient.BankAccountServiceClient.DeleteBankAccount(ctx, bankAccReq)
	err := bankAccountProto.ErrorToModel(protoErr)
	if err.IsNotNil() {
		return err
	}

	return errors.Nil()
}


func (si *ServiceImplementation) GetExpenses(ctx context.Context, bankAccountID string) ([]expense.Expense, errors.Error) {
	expensesReq := expenseProto.GetExpensesRequestToProto(bankAccountID)

	protoReply, _ := si.gRPCClient.ExpenseServiceClient.GetExpenses(ctx, expensesReq)

	err := expenseProto.ErrorToModel(protoReply.Error)
	if err.IsNotNil() {
		return nil, err
	}

	return expenseProto.GetExpensesReplyToModel(protoReply)
}

func (si *ServiceImplementation) CreateExpense(ctx context.Context, expense *expense.Expense) errors.Error {
	expenseReq := expenseProto.CreateExpenseRequestToProto(expense)

	protoErr, _ := si.gRPCClient.ExpenseServiceClient.CreateExpense(ctx, expenseReq)
	err := expenseProto.ErrorToModel(protoErr)
	if err.IsNotNil() {
		return err
	}

	return errors.Nil()
}

func (si *ServiceImplementation) DeleteExpense(ctx context.Context, id string) errors.Error {
	expenseReq := expenseProto.DeleteExpenseRequestToProto(id)

	protoErr, _ := si.gRPCClient.ExpenseServiceClient.DeleteExpense(ctx, expenseReq)
	err := expenseProto.ErrorToModel(protoErr)
	if err.IsNotNil() {
		return err
	}

	return errors.Nil()
}
