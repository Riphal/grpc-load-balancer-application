package middleware

import (
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/service/auth"
	"github.com/gin-gonic/gin"
)

// AuthorizeJWT validates the token from the http request, returning a 401 if it's not valid
func AuthorizeJWT(authService auth.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		token := authHeader[len("JWT "):]

		claims, err := authService.ValidateToken(ctx, token)
		if err.IsNotNil() {
			ctx.AbortWithStatusJSON(err.StatusCodeFromMap(), gin.H{
				"status": "fail",
				"error": gin.H{
					"message": err.Message,
					"type":    err.Type,
				},
			})

			return
		}

		ctx.Set("id", claims.ID)
		ctx.Set("email", claims.Email)

		ctx.Next()
	}
}
