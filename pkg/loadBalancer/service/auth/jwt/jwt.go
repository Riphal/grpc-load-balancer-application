package jwt

import (
	"fmt"
	"time"

	"github.com/Riphal/grpc-load-balancer-application/common/config"
	"github.com/Riphal/grpc-load-balancer-application/common/errors"
	"github.com/dgrijalva/jwt-go"
)

// jwtCustomClaims are custom claims extending default ones.
type jwtCustomClaims struct {
	ID  				string `json:"id"`
	Email  				string `json:"email"`
	jwt.StandardClaims
}

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
	claims := &jwtCustomClaims{
		id,
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
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

func (si *ServiceImplementation) ValidateToken(token string) errors.Error {
	key, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		// Signing method validation
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", jwtToken.Header["alg"])
		}

		// Return the secret signing key
		return []byte(si.secretKey), nil
	})

	if key != nil && key.Valid {
		return errors.Nil()
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return errors.New("that's not even a token", errors.ValidationError)
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			return errors.New("timing is everything", errors.ValidationError)
		}
	}

	return errors.New("couldn't handle this token", errors.BadRequestError)
}
