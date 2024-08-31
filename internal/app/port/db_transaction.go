package port

import (
	"context"
)

type dbTxMaker interface {
	BeginTx(ctx context.Context) (DBTx, error)
}

type DBTx interface {
	Commit() error
	Rollback() error
}

