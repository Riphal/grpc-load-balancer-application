package controller

import (
	"context"
	bankAccountProto "github.com/Riphal/grpc-load-balancer-application/common/proto/bankAccount"
	bankAccountService "github.com/Riphal/grpc-load-balancer-application/pkg/worker/service/bankAccount"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type BankAccountController struct {
	*Controller
	GRPCServer		*grpc.Server
	bankAccountService 	bankAccountService.Service
}

func NewBankAccountController(config *Config, bankAccountService bankAccountService.Service) *BankAccountController {
	return &BankAccountController{
		Controller:		NewController(config),
		bankAccountService:	bankAccountService,
	}
}

func (bac *BankAccountController) GetBankAccounts(ctx context.Context, empty *emptypb.Empty) (*bankAccountProto.BankAccountsReply, error) {
	panic("implement me")
}

func (bac *BankAccountController) GetBankAccount(ctx context.Context, request *bankAccountProto.BankAccountRequest) (*bankAccountProto.BankAccountReply, error) {
	panic("implement me")
}

func (bac *BankAccountController) CreateBankAccount(ctx context.Context, request *bankAccountProto.CreateBankAccountRequest) (*bankAccountProto.Error, error) {
	bankAcc := bankAccountProto.CreateBankAccountRequestToModel(request)

	return bac.bankAccountService.CreateBankAccount(ctx, bankAcc), nil
}

func (bac *BankAccountController) DeleteBankAccount(ctx context.Context, request *bankAccountProto.BankAccountRequest) (*bankAccountProto.Error, error) {
	panic("implement me")
}
