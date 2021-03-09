package response

type BankAccountsResponse struct {
	BankAccounts	[]BankAccountResponse `json:"bank_accounts"`
}

type BankAccountResponse struct {
	ID			string	`json:"id,omitempty" pg:",pk"`
	Name		string	`json:"name,omitempty" pg:",notnull"`
	Balance		float64	`json:"balance" pg:",notnull"`
}
