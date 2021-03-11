package bankAccount

import (
	"github.com/Riphal/grpc-load-balancer-application/common/errors"
	"github.com/Riphal/grpc-load-balancer-application/common/model/bankAccount"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/model/response"
)


func ErrorToProto(e errors.Error) *Error {
	return &Error{
		Message: 	e.Message,
		Type: 		e.Type,
	}
}

func ErrorToModel(e *Error) errors.Error {
	return errors.Error{
		Message: 	e.Message,
		Type: 		e.Type,
	}
}

func GetBankAccountsRequestToProto(accountID string) *BankAccountsRequest {
	return &BankAccountsRequest{
		AccountId: accountID,
	}
}

func GetBankAccountsReplyToProto(bankAccounts []*bankAccount.BankAccount, e errors.Error) *BankAccountsReply {
	var bankAccountsProto []*BankAccountReply

	for _, bankAcc := range bankAccounts {
		bankAccountsProto = append(bankAccountsProto, &BankAccountReply{
			Id:   bankAcc.ID,
			Name: bankAcc.Name,
		})
	}

	return &BankAccountsReply{
		BankAccounts: 	bankAccountsProto,
		Error: 			ErrorToProto(e),
	}
}

func GetBankAccountsReplyToModel(ba *BankAccountsReply) ([]response.BankAccountResponse, errors.Error) {
	var bankAccountsModel []response.BankAccountResponse

	for _, bankAcc := range ba.BankAccounts {
		bankAccountsModel = append(bankAccountsModel, response.BankAccountResponse{
			ID: 		bankAcc.Id,
			Name: 		bankAcc.Name,
			Balance:	bankAcc.Balance,
		})
	}

	return bankAccountsModel, ErrorToModel(ba.Error)
}

func GetBankAccountRequestToProto(id string) *BankAccountRequest {
	return &BankAccountRequest{
		Id: id,
	}
}

func GetBankAccountReplyToProto(ba *bankAccount.BankAccountBalance, e errors.Error) *BankAccountReply {
	return &BankAccountReply{
		Id: 		ba.ID,
		AccountId:  ba.AccountID,
		Name: 		ba.Name,
		Balance: 	ba.Balance,
		Error:		ErrorToProto(e),
	}
}

func GetBankAccountReplyToModel(ba *BankAccountReply) (*bankAccount.BankAccountBalance, errors.Error) {
	return &bankAccount.BankAccountBalance{
		ID: 		ba.Id,
		AccountID:  ba.AccountId,
		Name: 		ba.Name,
		Balance: 	ba.Balance,
	}, ErrorToModel(ba.Error)
}

func CreateBankAccountRequestToModel(ba *CreateBankAccountRequest) *bankAccount.BankAccount {
	return &bankAccount.BankAccount{
		AccountID:	 ba.AccountId,
		Name:		 ba.Name,
	}
}

func CreateBankAccountRequestToProto(ba *bankAccount.BankAccount) *CreateBankAccountRequest {
	return &CreateBankAccountRequest{
		AccountId:	ba.AccountID,
		Name: 		ba.Name,
	}
}
