package grpc

import (
	"context"
	"log"

	"github.com/Riphal/grpc-load-balancer-application/common/errors"
	"github.com/Riphal/grpc-load-balancer-application/common/model/bankAccount"
	"github.com/Riphal/grpc-load-balancer-application/common/model/expense"
	bankAccountProto "github.com/Riphal/grpc-load-balancer-application/common/proto/bankAccount"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/model/response"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/service"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/storage"
	"google.golang.org/grpc"
)

type Config struct {
	*service.Config
	LoadBalancerStorage	storage.LoadBalancer
}

type ServiceImplementation struct {
	*service.Service
	bankAccountServiceClient bankAccountProto.BankAccountServiceClient
	loadBalancerStorage	storage.LoadBalancer
}

func NewServiceImplementation(config *Config) *ServiceImplementation {
	return &ServiceImplementation{
		Service:        			service.New(config.Config),
		bankAccountServiceClient:	mustInitGRPCClientConn(),
		loadBalancerStorage:		config.LoadBalancerStorage,
	}
}

func mustInitGRPCClientConn() bankAccountProto.BankAccountServiceClient {
	conn, err := grpc.Dial(
		":9001",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("could't connect: %s", err)
	}

	return bankAccountProto.NewBankAccountServiceClient(conn)
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
	bankAccReq := bankAccountProto.CreateBankAccountRequestToProto(bankAccount)

	protoErr, _ := si.bankAccountServiceClient.CreateBankAccount(ctx, bankAccReq)
	err := bankAccountProto.ErrorToModel(protoErr)
	if err.IsNotNil() {
		return *err
	}

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
