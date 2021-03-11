package expense

import (
	"context"

	"github.com/Riphal/grpc-load-balancer-application/common/errors"
	"github.com/Riphal/grpc-load-balancer-application/common/model/expense"
	expenseProto "github.com/Riphal/grpc-load-balancer-application/common/proto/expense"
	"github.com/Riphal/grpc-load-balancer-application/pkg/worker/service"
	"github.com/Riphal/grpc-load-balancer-application/pkg/worker/storage"
)

type Config struct {
	*service.Config
	ExpenseStorage	storage.Expense
}

type ServiceImplementation struct {
	*service.Service
	expenseStorage	storage.Expense
}

func NewServiceImplementation(config *Config) *ServiceImplementation {
	return &ServiceImplementation{
		Service:			service.New(config.Config),
		expenseStorage:		config.ExpenseStorage,
	}
}


func (si ServiceImplementation) GetExpenses(ctx context.Context, bankAccountID string) *expenseProto.ExpensesReply {
	expenses, err := si.expenseStorage.GetExpenses(ctx, bankAccountID)
	if err.IsNotNil() {
		return &expenseProto.ExpensesReply{
			Error: expenseProto.ErrorToProto(err),
		}
	}

	return expenseProto.GetExpensesReplyToProto(expenses, err)
}

func (si ServiceImplementation) GetExpense(ctx context.Context, id string) *expenseProto.ExpenseReply {
	exp, err := si.expenseStorage.GetExpense(ctx, id)
	if err.IsNotNil() {
		return &expenseProto.ExpenseReply{
			Error: expenseProto.ErrorToProto(err),
		}
	}

	return expenseProto.GetExpenseReplyToProto(exp, err)
}

func (si ServiceImplementation) CreateExpense(ctx context.Context, expense *expense.Expense) *expenseProto.Error {
	err := si.expenseStorage.CreateExpense(ctx, expense)
	if err.IsNotNil() {
		return expenseProto.ErrorToProto(err)
	}

	return expenseProto.ErrorToProto(errors.Nil())
}

func (si ServiceImplementation) DeleteExpense(ctx context.Context, id string) *expenseProto.Error {
	err := si.expenseStorage.DeleteExpense(ctx, id)
	if err.IsNotNil() {
		return expenseProto.ErrorToProto(err)
	}

	return expenseProto.ErrorToProto(errors.Nil())
}
