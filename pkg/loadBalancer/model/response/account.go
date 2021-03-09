package response

type AccountResponse struct {
	Email 		string 	`json:"email,omitempty" pg:",notnull"`
	FirstName   string 	`json:"first_name,omitempty" pg:",use_zero"`
	LastName    string  `json:"last_name,omitempty" pg:",use_zero"`
}
