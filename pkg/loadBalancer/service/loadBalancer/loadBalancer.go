package loadBalancer

import (
	"context"
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/Riphal/grpc-load-balancer-application/common/errors"
	bankAccountProto "github.com/Riphal/grpc-load-balancer-application/common/proto/bankAccount"
	expenseProto "github.com/Riphal/grpc-load-balancer-application/common/proto/expense"
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
	loadBalancerStorage	storage.LoadBalancer
}

func NewServiceImplementation(config *Config) *ServiceImplementation {
	return &ServiceImplementation{
		Service:        		service.New(config.Config),
		loadBalancerStorage:	config.LoadBalancerStorage,
	}
}

var workers = map[string]*GRPCClient{}

func (si ServiceImplementation) RegisterWorker(ctx context.Context, addr string) errors.Error {
	conn, err := grpc.Dial(
		fmt.Sprintf(":%s", addr),
		grpc.WithInsecure(),
	)

	if err != nil {
		log.Fatalf("could't connect: %s", err)
	}

	workers[addr] = &GRPCClient{
		BankAccountServiceClient: 	bankAccountProto.NewBankAccountServiceClient(conn),
		ExpenseServiceClient:		expenseProto.NewExpenseServiceClient(conn),
	}

	return si.loadBalancerStorage.RegisterWorker(ctx, addr)
}

func (si ServiceImplementation) DeRegisterWorker(ctx context.Context, addr string) errors.Error {
	return si.loadBalancerStorage.DeRegisterWorker(ctx, addr)
}

func (si ServiceImplementation) GRPCClient(ctx context.Context) (*GRPCClient, errors.Error) {
	cursor, err := si.loadBalancerStorage.IncrCounter(ctx)
	if err.IsNotNil() {
		return nil, err
	}

	addrs, err := si.GetWorkerAddr(ctx)
	if err.IsNotNil() {
		return nil, err
	}

	// Round robin
	worker := workers[addrs[cursor % int64(len(addrs))]]

	return worker, err
}

func (si ServiceImplementation) GetWorkerAddr(ctx context.Context) ([]string, errors.Error) {
	var addr []string

	workers, err := si.loadBalancerStorage.GetWorkers(ctx)
	if err.IsNotNil() {
		return nil, err
	}

	for _, worker := range workers {
		splitter := strings.Split(worker, ":")

		addr = append(addr, splitter[1])
	}

	sort.Strings(addr)

	return addr, err
}
