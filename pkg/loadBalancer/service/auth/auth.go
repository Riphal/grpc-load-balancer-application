package auth

import (
	"context"
	"log"

	"github.com/Riphal/grpc-load-balancer-application/common/errors"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/model/account"
	jwtModel "github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/model/jwt"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/service"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/service/auth/jwt"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/storage"
	"golang.org/x/crypto/bcrypt"
)

type Config struct {
	*service.Config
	AuthStorage 	storage.Auth
	AccountStorage 	storage.Account
	JwtService		jwt.Service
}

type ServiceImplementation struct {
	*service.Service
	authStorage 	storage.Auth
	accountStorage 	storage.Account
	jwtService		jwt.Service
}

func NewServiceImplementation(config *Config) *ServiceImplementation {
	return &ServiceImplementation{
		Service:        service.New(config.Config),
		authStorage: 	config.AuthStorage,
		accountStorage: config.AccountStorage,
		jwtService:		config.JwtService,
	}
}

func (si *ServiceImplementation) Register(ctx context.Context, account *account.Account) (string, errors.Error) {
	// Check if bankAccount already exist
	_, err := si.accountStorage.GetAccountWithEmail(ctx, account.Email)
	if err.IsNil() {
		return "", errors.New("bankAccount already exists", errors.ConflictError)
	} else if err.IsNotNil() && err.Type != errors.PostgresNotFoundError {
		return "", err
	}

	// Save original password for login
	pass := account.Password

	// Hash password
	account.Password = hashAndSalt([]byte(account.Password))

	// Create bankAccount
	err = si.accountStorage.CreateAccount(ctx, account)
	if err.IsNotNil() {
		return "", err
	}

	// Call login to get valid token
	return si.Login(ctx, account.Email, pass)
}

func (si *ServiceImplementation) Login(ctx context.Context, email, password string) (string, errors.Error) {
	acc, err := si.accountStorage.GetAccountWithEmail(ctx, email)
	if err.IsNotNil() {
		return "", err
	}

	err = comparePasswords(acc.Password, []byte(password))
	if err.IsNotNil() {
		return "", err
	}

	return si.jwtService.GenerateToken(acc.ID, acc.Email)
}

func (si *ServiceImplementation) Logout(ctx context.Context, token string) errors.Error {
	err := si.authStorage.SetBlacklistToken(ctx, token)
	if err.IsNotNil() {
		return err
	}

	return errors.Nil()
}

func (si *ServiceImplementation) ValidateToken(ctx context.Context, token string) (*jwtModel.Claims, errors.Error) {
	claims, err := si.jwtService.ValidateToken(token)
	if err.IsNotNil() {
		return nil, err
	}

	exist, err := si.authStorage.IsBlacklisted(ctx, token)
	if err.IsNotNil() && err.Type != errors.RedisNotFoundError {
		return nil, err
	} else if exist {
		return nil, errors.New("this token is blacklisted", errors.ForbiddenError)
	}

	return claims, errors.Nil()
}

func hashAndSalt(pwd []byte) string {
	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd []byte) errors.Error {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return errors.New("incorrect password", errors.UnauthorizedError)
	}

	return errors.Nil()
}
