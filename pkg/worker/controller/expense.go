package controller

import (
	"context"
	expenseProto "github.com/Riphal/grpc-load-balancer-application/common/proto/expense"
	expenseService "github.com/Riphal/grpc-load-balancer-application/pkg/worker/service/expense"
)

type ExpenseController struct {
	*Controller
	expenseService 	expenseService.Service
}

func NewExpenseController(config *Config, expenseService expenseService.Service) *ExpenseController {
	return &ExpenseController{
		Controller:		NewController(config),
		expenseService:	expenseService,
	}
}

func (ec ExpenseController) GetExpenses(ctx context.Context, request *expenseProto.ExpensesRequest) (*expenseProto.ExpensesReply, error) {
	return ec.expenseService.GetExpenses(ctx, request.BankAccountId), nil
}

func (ec ExpenseController) CreateExpense(ctx context.Context, request *expenseProto.CreateExpenseRequest) (*expenseProto.Error, error) {
	exp := expenseProto.CreateExpenseRequestToModel(request)

	return ec.expenseService.CreateExpense(ctx, exp), nil
}

func (ec ExpenseController) DeleteExpense(ctx context.Context, request *expenseProto.DeleteExpenseRequest) (*expenseProto.Error, error) {
	return ec.expenseService.DeleteExpense(ctx, request.Id), nil
}
