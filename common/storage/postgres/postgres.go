package postgres

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Riphal/grpc-load-balancer-application/common/errors"
	"github.com/Riphal/grpc-load-balancer-application/common/storage"
	"github.com/go-pg/pg/v10"
	pkgerrors "github.com/pkg/errors"
)

type DBLogger struct{}

type DB struct {
	*pg.DB
}

var _ storage.Pinger = (*DB)(nil)

func New(url string) (*DB, error) {
	options, err := pg.ParseURL(url)
	if err != nil {
		return nil, pkgerrors.Wrap(err, fmt.Sprintf("error parsing connection url: %s", url))
	}

	log.Printf("options: %+v\n", options)

	db := pg.Connect(options)

	//db.AddQueryHook(&DBLogger{})

	err = db.Ping(context.Background())
	if err != nil {
		return nil, pkgerrors.Wrap(err, fmt.Sprintf("error connecting to postgres url: %s", url))
	}

	return &DB{ db }, nil
}

func (l *DBLogger) BeforeQuery(c context.Context, qe *pg.QueryEvent) (context.Context, error) {
	return c, nil
}
func (l *DBLogger) AfterQuery(c context.Context, qe *pg.QueryEvent) error {
	query, _ := qe.FormattedQuery()
	log.Println("executed query: ", string(query))

	return nil
}

func (db *DB) Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return db.Conn().Ping(ctx)
}

func (db *DB) HandleError(message string, err error) errors.Error {
	var errType string

	switch err {
	case pg.ErrNoRows:
		errType = errors.PostgresNotFoundError
	default:
		errType = errors.PostgresInternalError
	}

	return errors.New("pg: " + message, errType)
}
