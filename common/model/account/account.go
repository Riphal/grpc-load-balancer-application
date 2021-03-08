package account

type Account struct {
	ID			string	`json:"id,omitempty" pg:",pk"`
	Email 		string 	`json:"email,omitempty" pg:",notnull"`
	Password	string	`json:"password,omitempty" pg:",notnull"`
	FirstName   string 	`json:"first_name,omitempty" pg:",use_zero"`
	LastName    string  `json:"last_name,omitempty" pg:",use_zero"`
}
