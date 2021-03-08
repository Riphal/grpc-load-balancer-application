package account

import (
	"context"
	"fmt"

	"github.com/Riphal/grpc-load-balancer-application/common/errors"
	"github.com/Riphal/grpc-load-balancer-application/common/storage"
	"github.com/Riphal/grpc-load-balancer-application/common/storage/postgres"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/model/account"
)

type PGStorageImplementation struct {
	db *postgres.DB
}

var _ storage.Account = (*PGStorageImplementation)(nil)

func NewPGStorageImplementation(db *postgres.DB) *PGStorageImplementation {
	return &PGStorageImplementation{
		db: db,
	}
}

func (p *PGStorageImplementation) GetAccount(ctx context.Context, email string) (*account.Account, errors.Error) {
	acc := new(account.Account)

	err := p.db.ModelContext(ctx, acc).
		Where("email = ?", email).
		Select()

	if err != nil {
		return nil, p.db.HandleError(fmt.Sprintf("couldn't get account with email %s", email), err)
	}

	return acc, errors.Nil()
}

func (p *PGStorageImplementation) CreateAccount(ctx context.Context, account *account.Account) errors.Error {
	_, err := p.db.ModelContext(ctx, account).Insert()

	if err != nil {
		return p.db.HandleError("couldn't insert account", err)
	}

	return errors.Nil()
}

func (p *PGStorageImplementation) DeleteAccount(ctx context.Context, email string) errors.Error {
	acc := new(account.Account)

	_, err := p.db.ModelContext(ctx, acc).
		Where("email = ?", email).
		Delete()

	if err != nil {
		return p.db.HandleError(fmt.Sprintf("couldn't delete account with email %s", email), err)
	}

	return errors.Nil()
}
