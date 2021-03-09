package bankAccount

type BankAccount struct {
	ID			string	`json:"id,omitempty" pg:",pk"`
	AccountID 	string 	`json:"account_id,omitempty" pg:",notnull"`
	Name		string	`json:"name,omitempty" pg:",notnull"`
}
