package controller

import (
	"context"
	bankAccountProto "github.com/Riphal/grpc-load-balancer-application/common/proto/bankAccount"
	bankAccountService "github.com/Riphal/grpc-load-balancer-application/pkg/worker/service/bankAccount"
)

type BankAccountController struct {
	*Controller
	bankAccountService 	bankAccountService.Service
}

func NewBankAccountController(config *Config, bankAccountService bankAccountService.Service) *BankAccountController {
	return &BankAccountController{
		Controller:			NewController(config),
		bankAccountService:	bankAccountService,
	}
}

func (bac *BankAccountController) GetBankAccounts(ctx context.Context, request *bankAccountProto.BankAccountsRequest) (*bankAccountProto.BankAccountsReply, error) {
	return bac.bankAccountService.GetBankAccounts(ctx, request.AccountId), nil
}

func (bac *BankAccountController) GetBankAccount(ctx context.Context, request *bankAccountProto.BankAccountRequest) (*bankAccountProto.BankAccountReply, error) {
	return bac.bankAccountService.GetBankAccount(ctx, request.Id), nil
}

func (bac *BankAccountController) CreateBankAccount(ctx context.Context, request *bankAccountProto.CreateBankAccountRequest) (*bankAccountProto.Error, error) {
	bankAcc := bankAccountProto.CreateBankAccountRequestToModel(request)

	return bac.bankAccountService.CreateBankAccount(ctx, bankAcc), nil
}

func (bac *BankAccountController) DeleteBankAccount(ctx context.Context, request *bankAccountProto.BankAccountRequest) (*bankAccountProto.Error, error) {
	return bac.bankAccountService.DeleteBankAccount(ctx, request.Id), nil
}
