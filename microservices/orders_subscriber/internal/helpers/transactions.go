package helpers

import (
	"context"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"sync"
)

type TransactionsInterface interface {
	Rollback(ctx context.Context) error
	Commit(ctx context.Context) error
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

type Transactions struct {
	transactions pgx.Tx
	mutex        sync.Mutex
}

var _ TransactionsInterface = &Transactions{}

func NewTransactions(ctx context.Context, pool *pgxpool.Pool) (*Transactions, error) {
	transactions, err := pool.Begin(ctx)

	if err != nil {
		return nil, err
	}

	return &Transactions{transactions: transactions}, nil
}

func (t *Transactions) Rollback(ctx context.Context) error {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	return t.transactions.Rollback(ctx)
}

func (t *Transactions) Commit(ctx context.Context) error {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	return t.transactions.Commit(ctx)
}

func (t *Transactions) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	return t.transactions.QueryRow(ctx, sql, args...)
}

func (t *Transactions) Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	return t.transactions.Exec(ctx, sql, args...)
}
