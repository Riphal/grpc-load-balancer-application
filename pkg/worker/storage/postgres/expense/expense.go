package expense

import (
	"context"
	"fmt"
	"github.com/Riphal/grpc-load-balancer-application/common/errors"
	"github.com/Riphal/grpc-load-balancer-application/common/model/expense"
	"github.com/Riphal/grpc-load-balancer-application/common/storage/postgres"
	"github.com/Riphal/grpc-load-balancer-application/pkg/worker/storage"
)

type PGStorageImplementation struct {
	db *postgres.DB
}

var _ storage.Expense = (*PGStorageImplementation)(nil)

func NewPGStorageImplementation(db *postgres.DB) *PGStorageImplementation {
	return &PGStorageImplementation{
		db: db,
	}
}

func (p PGStorageImplementation) GetExpenses(ctx context.Context, bankAccountID string) ([]*expense.Expense, errors.Error) {
	var expenses []*expense.Expense

	err := p.db.ModelContext(ctx, &expenses).
		Where("bank_account_id = ?", bankAccountID).
		Select()

	if err != nil {
		return nil, p.db.HandleError(fmt.Sprintf("couldn't get expenses with bank account id %s", bankAccountID), err)
	}

	return expenses, errors.Nil()
}

func (p PGStorageImplementation) GetExpense(ctx context.Context, id string) (*expense.Expense, errors.Error) {
	exp := &expense.Expense{ ID: id }

	err := p.db.ModelContext(ctx, exp).WherePK().Select()
	if err != nil {
		return nil, p.db.HandleError(fmt.Sprintf("couldn't get expenese with id %s", id), err)
	}

	return exp, errors.Nil()
}

func (p PGStorageImplementation) CreateExpense(ctx context.Context, expense *expense.Expense) errors.Error {
	_, err := p.db.ModelContext(ctx, expense).Insert()
	if err != nil {
		return p.db.HandleError("couldn't insert expense", err)
	}

	return errors.Nil()
}

func (p PGStorageImplementation) DeleteExpense(ctx context.Context, id string) errors.Error {
	exp := &expense.Expense{ ID: id }

	_, err := p.db.ModelContext(ctx, exp).WherePK().Delete()
	if err != nil {
		return p.db.HandleError(fmt.Sprintf("couldn't delete expense with id %s", id), err)
	}

	return errors.Nil()
}
