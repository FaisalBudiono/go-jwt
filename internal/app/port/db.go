package port

import (
	"FaisalBudiono/go-jwt/internal/app/port/common"
	"context"
	"database/sql"
	"errors"
)

type DB interface {
	DB() *sql.DB
	FindUserByEmail(
		ctx context.Context, email string, tx *sql.Tx,
	) (common.User, error)
	InsertUser(
		ctx context.Context, u common.User, tx *sql.Tx,
	) (common.User, error)
}

var ErrDBResNotFound = errors.New("resource not found")
