package bankAccount

import "github.com/Riphal/grpc-load-balancer-application/common/model/expense"

type BankAccount struct {
	tableName 	struct{} 			`pg:"bank_accounts"`
	ID			string				`json:"id,omitempty" pg:",pk"`
	AccountID 	string 				`json:"account_id,omitempty" pg:",notnull"`
	Name		string				`json:"name,omitempty" pg:",notnull"`
	Expenses 	[]*expense.Expense 	`pg:"rel:has-many,join_fk:bank_account_id"`
}

type BankAccountBalance struct {
	tableName 	struct{} 			`pg:"bank_accounts"`
	ID			string				`json:"id,omitempty" pg:",pk"`
	AccountID 	string 				`json:"account_id,omitempty" pg:",notnull"`
	Name		string				`json:"name,omitempty" pg:",notnull"`
	Balance		float32				`json:"balance,omitempty"`
	Expenses 	[]*expense.Expense 	`pg:"rel:has-many,join_fk:bank_account_id"`
}
