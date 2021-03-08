package auth

import (
	"context"
	"fmt"

	"github.com/Riphal/grpc-load-balancer-application/common/errors"
	"github.com/Riphal/grpc-load-balancer-application/common/storage/postgres"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/model/auth"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/storage"
)

type PGStorageImplementation struct {
	db *postgres.DB
}

var _ storage.Auth = (*PGStorageImplementation)(nil)

func NewPGStorageImplementation(db *postgres.DB) *PGStorageImplementation {
	return &PGStorageImplementation{
		db: db,
	}
}

func (p *PGStorageImplementation) IsBlacklisted(ctx context.Context, token string) (bool, errors.Error) {
	ok, err := p.db.Model(&auth.Auth{ Token: token }).WherePK().Context(ctx).Exists()
	if err != nil {
		return true, p.db.HandleError(fmt.Sprintf("couldn't get token: %s", token), err)
	}

	return ok, errors.Nil()
}
