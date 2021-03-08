package auth

import (
	"context"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/service/auth/jwt"
	"golang.org/x/crypto/bcrypt"
	"log"

	"github.com/Riphal/grpc-load-balancer-application/common/errors"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/model/account"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/service"
	lbstorage "github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/storage"
)

type Config struct {
	*service.Config
	AuthStorage 	lbstorage.Auth
	AccountStorage 	lbstorage.Account
	JwtService		jwt.Service
}

type ServiceImplementation struct {
	*service.Service
	authStorage 	lbstorage.Auth
	accountStorage 	lbstorage.Account
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
	// Check if account already exist
	_, err := si.accountStorage.GetAccount(ctx, account.ID)
	if err.IsNil() {
		return "", errors.New("account already exists", errors.ConflictError)
	} else if err.IsNotNil() && err.Type != errors.PostgresNotFoundError {
		return "", err
	}

	// Hash password
	account.Password = hashAndSalt([]byte(account.Password))

	// Create account
	err = si.accountStorage.CreateAccount(ctx, account)
	if err.IsNotNil() {
		return "", err
	}

	// Call login to get valid token
	return si.Login(ctx, account)
}

func (si *ServiceImplementation) Login(ctx context.Context, account *account.Account) (string, errors.Error) {
	acc, err := si.accountStorage.GetAccount(ctx, account.Email)
	if err.IsNotNil() {
		return "", err
	}

	err = comparePasswords(acc.Password, []byte(account.Password))
	if err.IsNotNil() {
		return "", err
	}

	return si.jwtService.GenerateToken(account.Email)
}

func (si *ServiceImplementation) Logout(ctx context.Context, token string) errors.Error {
	err := si.authStorage.SetBlacklistToken(ctx, token)
	if err.IsNotNil() {
		return err
	}

	return errors.Nil()
}

func (si *ServiceImplementation) ValidateToken(ctx context.Context, token string) errors.Error {
	err := si.jwtService.ValidateToken(token)
	if err.IsNotNil() {
		return err
	}

	exist, err := si.authStorage.IsBlacklisted(ctx, token)
	if err.IsNotNil() && err.Type != errors.RedisNotFoundError {
		return err
	} else if exist {
		return errors.New("this token is blacklisted", errors.UnauthorizedError)
	}

	return errors.Nil()
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
