package response

type BankAccountResponse struct {
	ID			string	`json:"id,omitempty" pg:",pk"`
	Name		string	`json:"name,omitempty" pg:",notnull"`
	Balance		float32	`json:"balance,omitempty" pg:",notnull"`
}
