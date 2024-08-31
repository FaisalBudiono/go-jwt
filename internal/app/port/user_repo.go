package port

import (
	"FaisalBudiono/go-jwt/internal/app/domain"
	"context"
	"database/sql"
)

type UserRepo interface {
	InsertUser(ctx context.Context, u domain.User, tx *sql.Tx) (domain.User, error)
	FindUserByEmail(ctx context.Context, email string) (domain.User, error)
}
