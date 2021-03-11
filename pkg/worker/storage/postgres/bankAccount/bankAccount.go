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

func (p *PGStorageImplementation) GetBankAccounts(ctx context.Context, accountID string) ([]*bankAccount.BankAccount, errors.Error) {
	var bankAccounts []*bankAccount.BankAccount

	err := p.db.ModelContext(ctx, &bankAccounts).
		Where("account_id = ?", accountID).
		Select()

	if err != nil {
		return nil, p.db.HandleError(fmt.Sprintf("couldn't get bank accounts with account id %s", accountID), err)
	}

	return bankAccounts, errors.Nil()
}

func (p *PGStorageImplementation) GetBankAccount(ctx context.Context, id string) (*bankAccount.BankAccountBalance, errors.Error) {
	bankAcc := new(bankAccount.BankAccountBalance)

	_, err := p.db.QueryContext(ctx, bankAcc, `
		SELECT
			ba.id AS id,
			ba.account_id AS account_id,
			ba.name AS name,
			SUM(e.amount) AS balance
		FROM bank_accounts ba
		LEFT JOIN expenses e on ba.id = e.bank_account_id
		WHERE ba.id = ?
		GROUP BY ba.id, ba.account_id, ba.name;
	`, id)


	if err != nil {
		return nil, p.db.HandleError(fmt.Sprintf("couldn't get bank account with id %s", id), err)
	} else if bankAcc.ID == "" {
		return nil, errors.New(fmt.Sprintf("couldn't get bank account with id %s", id), errors.PostgresNotFoundError)
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

func (p *PGStorageImplementation) DeleteBankAccount(ctx context.Context, id string) errors.Error {
	bankAcc := &bankAccount.BankAccount{ ID: id }

	_, err := p.db.ModelContext(ctx, bankAcc).WherePK().Delete()
	if err != nil {
		return p.db.HandleError(fmt.Sprintf("couldn't delete bank account with id %s", id), err)
	}

	return errors.Nil()
}
