package pgdb

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type txCtxKey struct{}

func NewDatabase(pool *pgxpool.Pool) *Database {
	return &Database{
		pool: pool,
	}
}

type Database struct {
	pool *pgxpool.Pool
}

func (d *Database) QuerySq(ctx context.Context, query sq.Sqlizer) (pgx.Rows, error) {
	tx, withTransaction := TransactionFromContext(ctx)

	querySql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	if withTransaction {
		return tx.Query(ctx, querySql, args...)
	}
	return d.pool.Query(ctx, querySql, args...)
}

func TransactionFromContext(ctx context.Context) (pgx.Tx, bool) {
	if tx := ctx.Value(txCtxKey{}); tx != nil {
		return tx.(pgx.Tx), true
	}
	return nil, false
}
