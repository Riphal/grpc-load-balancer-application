package expense

import (
	"github.com/Riphal/grpc-load-balancer-application/common/errors"
	"github.com/Riphal/grpc-load-balancer-application/common/model/expense"
)


func ErrorToProto(e errors.Error) *Error {
	return &Error{
		Message: 	e.Message,
		Type: 		e.Type,
	}
}

func ErrorToModel(e *Error) errors.Error {
	return errors.Error{
		Message: 	e.Message,
		Type: 		e.Type,
	}
}

func GetExpensesRequestToProto(bankAccountID string) *ExpensesRequest {
	return &ExpensesRequest{
		BankAccountId: bankAccountID,
	}
}

func GetExpensesReplyToProto(expenses []*expense.Expense, e errors.Error) *ExpensesReply {
	var expensesProto []*Expense

	for _, exp := range expenses {
		expensesProto = append(expensesProto, &Expense{
			Id:   	exp.ID,
			Name: 	exp.Name,
			Amount: exp.Amount,
		})
	}

	return &ExpensesReply{
		Expenses: 	expensesProto,
		Error: 		ErrorToProto(e),
	}
}

func GetExpensesReplyToModel(replay *ExpensesReply) ([]expense.Expense, errors.Error) {
	var expenses []expense.Expense

	for _, exp := range replay.Expenses {
		expenses = append(expenses, expense.Expense{
			ID:   	exp.Id,
			Name: 	exp.Name,
			Amount: exp.Amount,
		})
	}

	return expenses, ErrorToModel(replay.Error)
}

func CreateExpenseRequestToModel(e *CreateExpenseRequest) *expense.Expense {
	return &expense.Expense{
		BankAccountID:	e.BankAccountId,
		Name: 			e.Name,
		Amount: 		e.Amount,
	}
}

func CreateExpenseRequestToProto(exp *expense.Expense) *CreateExpenseRequest {
	return &CreateExpenseRequest{
		BankAccountId: 	exp.BankAccountID,
		Name: 			exp.Name,
		Amount:			exp.Amount,
	}
}

func DeleteExpenseRequestToProto(id string) *DeleteExpenseRequest {
	return &DeleteExpenseRequest{
		Id: id,
	}
}
