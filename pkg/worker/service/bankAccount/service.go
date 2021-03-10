package bankAccount

import (
	"context"

	"github.com/Riphal/grpc-load-balancer-application/common/model/bankAccount"
	bankAccountProto "github.com/Riphal/grpc-load-balancer-application/common/proto/bankAccount"
)

type Service interface {
	GetBankAccounts(ctx context.Context, accountID string) *bankAccountProto.BankAccountsReply
	GetBankAccount(ctx context.Context, id string) *bankAccountProto.BankAccountReply
	CreateBankAccount(ctx context.Context, bankAccount *bankAccount.BankAccount) *bankAccountProto.Error
	DeleteBankAccount(ctx context.Context, id string) *bankAccountProto.Error
}
