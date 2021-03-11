package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	core "github.com/Riphal/grpc-load-balancer-application"
	bankAccountProto "github.com/Riphal/grpc-load-balancer-application/common/proto/bankAccount"
	expenseProto "github.com/Riphal/grpc-load-balancer-application/common/proto/expense"
	loadBalancerProto "github.com/Riphal/grpc-load-balancer-application/common/proto/loadBalancer"
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
	port := os.Getenv("PORT")

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


	// Register services to gRPC server
	bankAccountProto.RegisterBankAccountServiceServer(grpcServer, bankAccountController)
	expenseProto.RegisterExpenseServiceServer(grpcServer, expenseController)


	go heartBeatLoadBalancer(port)
	go deRegisterWorker(port)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("faild to serve gRPC server over %v: %v", port, err)
	}
}

func mustInitGRPC () (net.Listener, *grpc.Server) {
	port := os.Getenv("PORT")

	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalf("faild to listen on port %v: %v", port, err)
	}

	grpcServer := grpc.NewServer()

	log.Println(fmt.Sprintf("🚀 worker gRPC server listen on %v", listener.Addr()))

	return listener, grpcServer
}


func heartBeatLoadBalancer(addr string) {
	registerWorker(addr)

	for range time.Tick(10 * time.Second) {
		registerWorker(addr)
	}
}

func registerWorker(addr string) {
	conn, err := grpc.Dial(
		":4090",
		grpc.WithInsecure(),
	)

	if err != nil {
		return
	}
	defer conn.Close()

	wsClient := loadBalancerProto.NewLoadBalancerServiceClient(conn)

	_, err = wsClient.Register(context.Background(), loadBalancerProto.RequestToProto(addr))
}

func deRegisterWorker(addr string) {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs

	conn, err := grpc.Dial(
		":4090",
		grpc.WithInsecure(),
	)

	if err != nil {
		os.Exit(1)
	}
	defer conn.Close()

	wsClient := loadBalancerProto.NewLoadBalancerServiceClient(conn)

	_, err = wsClient.DeRegister(context.Background(), loadBalancerProto.RequestToProto(addr))

	os.Exit(1)
}
