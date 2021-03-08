package storage

import (
	"context"

	"github.com/Riphal/grpc-load-balancer-application/common/errors"
)

type Auth interface {
	IsBlacklisted(ctx context.Context, token string) (bool, errors.Error)
}
