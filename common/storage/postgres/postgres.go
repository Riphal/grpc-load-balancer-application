package postgres

import (
	"context"
	"fmt"
	"log"

	"github.com/Riphal/grpc-load-balancer-application/common/errors"
	"github.com/go-pg/pg/v10"
)

type DBLogger struct{}

type DB struct {
	*pg.DB
}

func New(url string) (*DB, errors.Error) {
	options, err := pg.ParseURL(url)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error parsing postgres url: %s", url), errors.InternalServerError)
	}

	db := pg.Connect(options)

	//db.AddQueryHook(&DBLogger{})

	err = db.Ping(context.Background())
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error connecting to postgres url: %s", url), errors.InternalServerError)
	}

	return &DB{ db }, errors.Nil()
}

func (l *DBLogger) BeforeQuery(c context.Context, qe *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

func (l *DBLogger) AfterQuery(c context.Context, qe *pg.QueryEvent) error {
	query, _ := qe.FormattedQuery()
	log.Println("executed query: ", string(query))

	return nil
}

func (db *DB) HandleError(message string, err error) errors.Error {
	var errType string

	switch err {
	case pg.ErrNoRows:
		errType = errors.PostgresNotFoundError
	default:
		errType = errors.PostgresInternalError
	}

	return errors.New(fmt.Sprintf("pg: %s", message), errType)
}
