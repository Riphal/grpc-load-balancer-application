package expense

type Expense struct {
	tableName 		struct{} 	`pg:"expenses"`
	ID				string		`json:"id,omitempty" pg:",pk"`
	BankAccountID 	string 		`json:"bank_account_id,omitempty" pg:",notnull"`
	Name			string		`json:"name,omitempty" pg:",notnull"`
	Amount			float32		`json:"amount,omitempty" pg:",notnull"`
}
