package controller

import (
	"net/http"

	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/model/account"
	authService "github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/service/auth"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	*Controller
	authService authService.Service
}

func NewAuthController(config *Config, authService authService.Service) *AuthController {
	return &AuthController{
		Controller:		NewController(config),
		authService:	authService,
	}
}

func (ac *AuthController) Register(ctx *gin.Context) {
	acc := &account.Account{}

	err := ac.BindRequestBodyAndHandleError(ctx, acc)
	if err.IsNotNil() {
		return
	}

	token, err := ac.authService.Register(ctx.Request.Context(), acc)
	if err.IsNotNil() {
		ac.AbortWithError(ctx, err)
		return
	}

	ac.JSON(ctx, http.StatusOK, gin.H{
		"status":	"ok",
		"token": 	token,
	})
}

func (ac *AuthController) Login(ctx *gin.Context) {
	acc := &account.Account{}

	err := ac.BindRequestBodyAndHandleError(ctx, acc)
	if err.IsNotNil() {
		return
	}

	token, err := ac.authService.Login(ctx.Request.Context(), acc.Email, acc.Password)
	if err.IsNotNil() {
		ac.AbortWithError(ctx, err)
		return
	}

	ac.JSON(ctx, http.StatusOK, gin.H{
		"status":	"ok",
		"token": 	token,
	})
}

func (ac *AuthController) Logout(ctx *gin.Context) {
	token := getTokenFromHeaders(ctx)

	err := ac.authService.Logout(ctx.Request.Context(), token)
	if err.IsNotNil() {
		ac.AbortWithError(ctx, err)
		return
	}

	ac.JSON(ctx, http.StatusOK, gin.H{
		"status":	"ok",
	})
}

func getTokenFromHeaders(ctx *gin.Context) string {
	authHeader := ctx.GetHeader("Authorization")
	tokenString := authHeader[len("JWT "):]

	return tokenString
}
