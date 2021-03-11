package expense

import (
	"context"
	"github.com/Riphal/grpc-load-balancer-application/common/model/expense"
	expenseProto "github.com/Riphal/grpc-load-balancer-application/common/proto/expense"
)

type Service interface {
	GetExpenses(ctx context.Context, bankAccountID string) *expenseProto.ExpensesReply
	GetExpense(ctx context.Context, id string) *expenseProto.ExpenseReply
	CreateExpense(ctx context.Context, expense *expense.Expense) *expenseProto.Error
	DeleteExpense(ctx context.Context, id string) *expenseProto.Error
}
