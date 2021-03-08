package account

import (
	"context"
	"fmt"

	"github.com/Riphal/grpc-load-balancer-application/common/errors"
	"github.com/Riphal/grpc-load-balancer-application/common/model/account"
	"github.com/Riphal/grpc-load-balancer-application/common/storage"
	"github.com/Riphal/grpc-load-balancer-application/common/storage/postgres"
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

func (p *PGStorageImplementation) GetAccount(ctx context.Context, id string) (*account.Account, errors.Error) {
	acc := &account.Account{ ID: id }

	err := p.db.Model(acc).WherePK().Context(ctx).Select()
	if err != nil {
		return nil, p.db.HandleError(fmt.Sprintf("couldn't get accounts with id %s", id), err)
	}

	return acc, errors.Nil()
}

func (p *PGStorageImplementation) CreateAccount(ctx context.Context, account *account.Account) errors.Error {
	_, err := p.db.Model(account).Context(ctx).Insert()

	if err != nil {
		return p.db.HandleError("couldn't insert account", err)
	}

	return errors.Nil()
}

func (p *PGStorageImplementation) DeleteAccount(ctx context.Context, id string) errors.Error {
	acc := &account.Account{ ID: id }

	_, err := p.db.Model(acc).WherePK().Delete()
	if err != nil {
		return p.db.HandleError(fmt.Sprintf("couldn't delete accounts with id %s", id), err)
	}

	return errors.Nil()
}

