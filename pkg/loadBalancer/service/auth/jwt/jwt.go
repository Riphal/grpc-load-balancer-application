package jwt

import (
	"fmt"
	"time"

	"github.com/Riphal/grpc-load-balancer-application/common/config"
	"github.com/Riphal/grpc-load-balancer-application/common/errors"
	jwtModel "github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/model/jwt"
	"github.com/dgrijalva/jwt-go"
)

type ServiceImplementation struct {
	secretKey string
}

func NewServiceImplementation() *ServiceImplementation {
	return &ServiceImplementation{
		secretKey: config.GetEnv("JWT_SECRET", "secret"),
	}
}

func (si *ServiceImplementation) GenerateToken(id, email string) (string, errors.Error) {
	// Set custom and standard claims
	claims := &jwtModel.Claims{
		ID:    	id,
		Email: 	email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: 	time.Now().Add(24 * time.Hour).Unix(),
			IssuedAt: 	time.Now().Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token using the secret signing key
	t, err := token.SignedString([]byte(si.secretKey))
	if err != nil {
		return "", errors.New("failed on generate token", errors.JWTGenerateTokenError)
	}

	return t, errors.Nil()
}

func (si *ServiceImplementation) ValidateToken(signedToken string) (*jwtModel.Claims, errors.Error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&jwtModel.Claims{},
		func(jwtToken *jwt.Token) (interface{}, error) {
		// Signing method validation
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", jwtToken.Header["alg"])
		}

		// Return the secret signing key
		return []byte(si.secretKey), nil
	})

	if token != nil && token.Valid {
		claims, ok := token.Claims.(*jwtModel.Claims)
		if !ok {
			return nil, errors.New("couldn't parse claims", errors.BadRequestError)
		}

		return claims, errors.Nil()
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return nil, errors.New("that's not even a token", errors.ValidationError)
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			return nil, errors.New("timing is everything", errors.ValidationError)
		}
	}

	return nil, errors.New("couldn't handle this token", errors.BadRequestError)
}
