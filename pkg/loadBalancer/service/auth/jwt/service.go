package jwt

import (
	"github.com/Riphal/grpc-load-balancer-application/common/errors"
)

type Service interface {
	GenerateToken(id, email string) (string, errors.Error)
	ValidateToken(token string) errors.Error
}
