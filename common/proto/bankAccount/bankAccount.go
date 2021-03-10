package bankAccount

import (
	"github.com/Riphal/grpc-load-balancer-application/common/errors"
	"github.com/Riphal/grpc-load-balancer-application/common/model/bankAccount"
)

func ErrorToModel(e *Error) *errors.Error {
	return &errors.Error{
		Message: e.Message,
		Type: e.Type,
	}
}

func ErrorToProto(e *errors.Error) *Error {
	return &Error{
		Message: e.Message,
		Type: e.Type,
	}
}

func NilErrorToProto() *Error {
	return &Error{
		Message: "",
		Type: "",
	}
}

func CreateBankAccountRequestToModel(ba *CreateBankAccountRequest) *bankAccount.BankAccount {
	return &bankAccount.BankAccount{
		AccountID: ba.AccountId,
		Name: ba.Name,
	}
}

func CreateBankAccountRequestToProto(ba *bankAccount.BankAccount) *CreateBankAccountRequest {
	return &CreateBankAccountRequest{
		AccountId: ba.AccountID,
		Name: ba.Name,
	}
}
