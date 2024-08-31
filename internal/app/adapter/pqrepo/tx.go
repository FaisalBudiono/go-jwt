package pqrepo

import (
	"database/sql"
)

func newTx(tx *sql.Tx) *sqlTx {
	return &sqlTx{
		tx: tx,
	}
}

type sqlTx struct {
	tx *sql.Tx
}

func (t *sqlTx) Commit() error {
	return t.tx.Commit()
}

func (t *sqlTx) Rollback() error {
	return t.tx.Rollback()
}
