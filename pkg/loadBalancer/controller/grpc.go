package controller

import (
	"net/http"

	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/model/bankAccount"
	grpcService "github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/service/grpc"
	"github.com/gin-gonic/gin"
)

type GRPCController struct {
	*Controller
	grpcService grpcService.Service
}

func NewGRPCController(config *Config, grpcService grpcService.Service) *GRPCController {
	return &GRPCController{
		Controller:		NewController(config),
		grpcService:	grpcService,
	}
}


func (ac *GRPCController) GetAccount(ctx *gin.Context) {
	// TO-DO: Implement
}


func (ac *GRPCController) GetBankAccounts(ctx *gin.Context) {
	// TO-DO: Implement
}

func (ac *GRPCController) GetBankAccount(ctx *gin.Context) {
	// TO-DO: Implement
}

func (ac *GRPCController) CreateBankAccount(ctx *gin.Context) {
	bankAcc := &bankAccount.BankAccount{}

	err := ac.BindRequestBodyAndHandleError(ctx, bankAcc)
	if err.IsNotNil() {
		return
	}

	err = ac.grpcService.CreateBankAccount(ctx.Request.Context(), bankAcc)
	if err.IsNotNil() {
		ac.AbortWithError(ctx, err)
		return
	}

	ac.JSON(ctx, http.StatusOK, gin.H{
		"status":	"ok",
	})
}

func (ac *GRPCController) DeleteBankAccount(ctx *gin.Context) {
	// TO-DO: Implement
}


func (ac *GRPCController) GetExpenses(ctx *gin.Context) {
	// TO-DO: Implement
}

func (ac *GRPCController) CreateExpense(ctx *gin.Context) {
	// TO-DO: Implement
}

func (ac *GRPCController) DeleteExpense(ctx *gin.Context) {
	// TO-DO: Implement
}
