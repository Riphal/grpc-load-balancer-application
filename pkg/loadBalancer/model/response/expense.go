package response

import "github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/model/expense"

type ExpensesResponse struct {
	Expenses	[]expense.Expense `json:"expenses"`
}
