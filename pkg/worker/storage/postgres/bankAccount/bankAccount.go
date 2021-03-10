package bankAccount

import (
	"context"
	"fmt"

	"github.com/Riphal/grpc-load-balancer-application/common/errors"
	"github.com/Riphal/grpc-load-balancer-application/common/model/bankAccount"
	"github.com/Riphal/grpc-load-balancer-application/common/storage/postgres"
	"github.com/Riphal/grpc-load-balancer-application/pkg/worker/storage"
)

type PGStorageImplementation struct {
	db *postgres.DB
}

var _ storage.BankAccount = (*PGStorageImplementation)(nil)

func NewPGStorageImplementation(db *postgres.DB) *PGStorageImplementation {
	return &PGStorageImplementation{
		db: db,
	}
}

func (p *PGStorageImplementation) GetBankAccount(ctx context.Context, id string) (*bankAccount.BankAccount, errors.Error) {
	bankAcc := &bankAccount.BankAccount{ ID: id }

	err := p.db.ModelContext(ctx, bankAcc).WherePK().Select()
	if err != nil {
		return nil, p.db.HandleError(fmt.Sprintf("couldn't get bank account with id %s", id), err)
	}

	return bankAcc, errors.Nil()
}

func (p *PGStorageImplementation) CreateBankAccount(ctx context.Context, bankAccount *bankAccount.BankAccount) errors.Error {
	_, err := p.db.ModelContext(ctx, bankAccount).Insert()
	if err != nil {
		return p.db.HandleError("couldn't insert bank account", err)
	}

	return errors.Nil()
}
