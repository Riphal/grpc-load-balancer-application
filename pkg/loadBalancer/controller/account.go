package controller

import (
	accountService "github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/service/account"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccountController struct {
	*Controller
	accountService accountService.Service
}

func NewAccountController(config *Config, accountService accountService.Service) *AccountController {
	return &AccountController{
		Controller:		NewController(config),
		accountService:	accountService,
	}
}

func (ac *AccountController) GetAccount(ctx *gin.Context) {
	id, _ := ctx.Get("id")

	acc, err := ac.accountService.GetAccount(ctx, id.(string))
	if err.IsNotNil() {
		return
	}

	ac.JSON(ctx, http.StatusOK, gin.H{
		"status":	"ok",
		"account":	acc,
	})
}
