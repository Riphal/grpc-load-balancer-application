package controller

import (
	"net/http"

	"github.com/Riphal/grpc-load-balancer-application/common/model/bankAccount"
	"github.com/Riphal/grpc-load-balancer-application/common/model/expense"
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
	accountID, _ := ctx.Get("id")

	bankAccounts, err := c.grpcService.GetBankAccounts(ctx.Request.Context(), accountID.(string))
	if err.IsNotNil() {
		c.AbortWithError(ctx, err)
		return
	}

	c.JSON(ctx, http.StatusOK, gin.H{
		"status":			"ok",
		"bank_accounts": 	bankAccounts,
	})
}

func (c *GRPCController) GetBankAccount(ctx *gin.Context) {
	bankAccountID, _ := ctx.Params.Get("bank_account_id")

	bankAcc, err := c.grpcService.GetBankAccount(ctx.Request.Context(), bankAccountID)
	if err.IsNotNil() {
		c.AbortWithError(ctx, err)
		return
	}

	c.JSON(ctx, http.StatusOK, gin.H{
		"status":		"ok",
		"bank_account":	bankAcc,
	})
}

func (c *GRPCController) CreateBankAccount(ctx *gin.Context) {
	accountID, _ := ctx.Get("id")

	bankAcc := new(bankAccount.BankAccount)

	err := c.BindRequestBodyAndHandleError(ctx, bankAcc)
	if err.IsNotNil() {
		return
	}

	bankAcc.AccountID = accountID.(string)

	err = c.grpcService.CreateBankAccount(ctx.Request.Context(), bankAcc)
	if err.IsNotNil() {
		c.AbortWithError(ctx, err)
		return
	}

	c.JSON(ctx, http.StatusOK, gin.H{
		"status": "ok",
	})
}

func (c *GRPCController) DeleteBankAccount(ctx *gin.Context) {
	bankAccountID, _ := ctx.Params.Get("bank_account_id")

	err := c.grpcService.DeleteBankAccount(ctx.Request.Context(), bankAccountID)
	if err.IsNotNil() {
		c.AbortWithError(ctx, err)
		return
	}

	c.JSON(ctx, http.StatusOK, gin.H{
		"status": "ok",
	})
}


func (c *GRPCController) GetExpenses(ctx *gin.Context) {
	bankAccountID, _ := ctx.Params.Get("bank_account_id")

	expenses, err := c.grpcService.GetExpenses(ctx.Request.Context(), bankAccountID)
	if err.IsNotNil() {
		c.AbortWithError(ctx, err)
		return
	}

	c.JSON(ctx, http.StatusOK, gin.H{
		"status":	"ok",
		"expenses": expenses,
	})
}

func (c *GRPCController) CreateExpense(ctx *gin.Context) {
	bankAccountID, _ := ctx.Params.Get("bank_account_id")

	exp := new(expense.Expense)

	err := c.BindRequestBodyAndHandleError(ctx, exp)
	if err.IsNotNil() {
		return
	}

	exp.BankAccountID = bankAccountID

	err = c.grpcService.CreateExpense(ctx.Request.Context(), exp)
	if err.IsNotNil() {
		c.AbortWithError(ctx, err)
		return
	}

	c.JSON(ctx, http.StatusOK, gin.H{
		"status": "ok",
	})
}

func (c *GRPCController) DeleteExpense(ctx *gin.Context) {
	expenseID, _ := ctx.Params.Get("expense_id")

	err := c.grpcService.DeleteExpense(ctx.Request.Context(), expenseID)
	if err.IsNotNil() {
		c.AbortWithError(ctx, err)
		return
	}

	c.JSON(ctx, http.StatusOK, gin.H{
		"status": "ok",
	})
}
