package jwt

import (
	"github.com/Riphal/grpc-load-balancer-application/common/errors"
	jwtModel "github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/model/jwt"
)

type Service interface {
	GenerateToken(id, email string) (string, errors.Error)
	ValidateToken(token string) (*jwtModel.Claims, errors.Error)
}
