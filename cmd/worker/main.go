package main

import (
	"fmt"
	"log"
	"net"
	"os"

	core "github.com/Riphal/grpc-load-balancer-application"
	bankAccountProto "github.com/Riphal/grpc-load-balancer-application/common/proto/bankAccount"
	expenseProto "github.com/Riphal/grpc-load-balancer-application/common/proto/expense"
	"github.com/Riphal/grpc-load-balancer-application/pkg/worker/controller"
	"github.com/Riphal/grpc-load-balancer-application/pkg/worker/service"
	"github.com/Riphal/grpc-load-balancer-application/pkg/worker/service/bankAccount"
	"github.com/Riphal/grpc-load-balancer-application/pkg/worker/service/expense"
	bankAccountStorage "github.com/Riphal/grpc-load-balancer-application/pkg/worker/storage/postgres/bankAccount"
	expenseStorage "github.com/Riphal/grpc-load-balancer-application/pkg/worker/storage/postgres/expense"
	"google.golang.org/grpc"
)

func main() {
	app := core.NewApp()
	listener, grpcServer := mustInitGRPC()

	serviceConfig := &service.Config{}
	controllerConfig := &controller.Config{}


	// Init Services
	var bankAccountService bankAccount.Service = bankAccount.NewServiceImplementation(&bankAccount.Config{
		Config: 			serviceConfig,
		BankAccountStorage:	bankAccountStorage.NewPGStorageImplementation(app.DB),
	})

	var expenseService expense.Service = expense.NewServiceImplementation(&expense.Config{
		Config: 			serviceConfig,
		ExpenseStorage:		expenseStorage.NewPGStorageImplementation(app.DB),
	})


	// Init controllers
	bankAccountController := controller.NewBankAccountController(controllerConfig, bankAccountService)
	expenseController := controller.NewExpenseController(controllerConfig, expenseService)


	// register services to gRPC server
	bankAccountProto.RegisterBankAccountServiceServer(grpcServer, bankAccountController)
	expenseProto.RegisterExpenseServiceServer(grpcServer, expenseController)


	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("faild to serve gRPC server over %v: %v", listener.Addr(), err)
	}
}

func mustInitGRPC () (net.Listener, *grpc.Server) {
	port := os.Getenv("PORT")

	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalf("faild to listen on port %v: %v", port, err)
	}

	grpcServer := grpc.NewServer()

	log.Println(fmt.Sprintf("🚀 worker listen on %v", listener.Addr()))

	return listener, grpcServer
}
