package controller

import (
	"github.com/Riphal/grpc-load-balancer-application/common/errors"
	"github.com/gin-gonic/gin"
)

type Config struct {}

type Controller struct {}

func NewController(config *Config) *Controller {
	return &Controller{}
}

func (c *Controller) AbortWithError(ctx *gin.Context, err errors.Error) {
	ctx.AbortWithStatusJSON(err.StatusCodeFromMap(), gin.H{
		"status": "fail",
		"error": gin.H{
			"message": err.Message,
			"type":    err.Type,
		},
	})
}

func (c *Controller) JSON(ctx *gin.Context, code int, response interface{}) {
	ctx.JSON(code, response)
}

func (c *Controller) BindRequestBodyAndHandleError(ctx *gin.Context, requestBody interface{}) errors.Error {
	err := ctx.BindJSON(requestBody)
	if err != nil {
		modelErr := errors.New("failed unmarshalling request body", errors.ValidationError)

		c.AbortWithError(ctx, modelErr)

		return modelErr
	}

	return errors.Nil()
}

func (c *Controller) BadRequest(ctx *gin.Context, message string) {
	c.AbortWithError(ctx, errors.New(message, errors.BadRequestError))
}
