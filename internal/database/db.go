package database

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type DBTX interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

type Queries struct {
	DB DBTX
}

func New(db DBTX) *Queries {
	return &Queries{DB: db}
}

func (q *Queries) WithTx(tx pgx.Tx) *Queries {
	return &Queries{
		DB: tx,
	}
}
