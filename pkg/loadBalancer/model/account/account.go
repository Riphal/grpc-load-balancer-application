package account

import "github.com/Riphal/grpc-load-balancer-application/common/model/bankAccount"

type Account struct {
	tableName	 struct{} 					`pg:"accounts"`
	ID			 string						`json:"id,omitempty" pg:",pk"`
	Email 		 string 					`json:"email,omitempty" pg:",notnull"`
	Password	 string						`json:"password,omitempty" pg:",notnull"`
	FirstName	 string 					`json:"first_name,omitempty" pg:",use_zero"`
	LastName   	 string 	 				`json:"last_name,omitempty" pg:",use_zero"`
	BankAccounts []*bankAccount.BankAccount `pg:"rel:has-many,join_fk:account_id"`
}
