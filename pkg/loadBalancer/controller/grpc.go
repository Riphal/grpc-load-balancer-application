package controller

import (
	"net/http"

	"github.com/Riphal/grpc-load-balancer-application/common/model/bankAccount"
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


func (c *GRPCController) GetBankAccounts(ctx *gin.Context) {
	// TO-DO: Implement
}

func (c *GRPCController) GetBankAccount(ctx *gin.Context) {
	// TO-DO: Implement
}

func (c *GRPCController) CreateBankAccount(ctx *gin.Context) {
	id, _ := ctx.Get("id")

	bankAcc := &bankAccount.BankAccount{}

	err := c.BindRequestBodyAndHandleError(ctx, bankAcc)
	if err.IsNotNil() {
		return
	}

	bankAcc.AccountID = id.(string)

	err = c.grpcService.CreateBankAccount(ctx.Request.Context(), bankAcc)
	if err.IsNotNil() {
		c.AbortWithError(ctx, err)
		return
	}

	c.JSON(ctx, http.StatusOK, gin.H{
		"status":	"ok",
	})
}

func (c *GRPCController) DeleteBankAccount(ctx *gin.Context) {
	// TO-DO: Implement
}


func (c *GRPCController) GetExpenses(ctx *gin.Context) {
	// TO-DO: Implement
}

func (c *GRPCController) CreateExpense(ctx *gin.Context) {
	// TO-DO: Implement
}

func (c *GRPCController) DeleteExpense(ctx *gin.Context) {
	// TO-DO: Implement
}
