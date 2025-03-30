package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
)

// TxFn represents a function that will be executed within a transaction
type TxFn func(ctx context.Context, tx *sqlx.Tx) error

// TransactionManager defines the interface for transaction management
type TransactionManager interface {
	// WithTransaction executes the given function within a transaction
	WithTransaction(ctx context.Context, fn TxFn) error
}
