package expense

type Expense struct {
	ID				string	`json:"id,omitempty" pg:",pk"`
	BankAccountID 	string 	`json:"bank_account_id,omitempty" pg:",notnull"`
	Name			string	`json:"name,omitempty" pg:",notnull"`
	Amount			string	`json:"amount,omitempty" pg:",notnull"`
}
