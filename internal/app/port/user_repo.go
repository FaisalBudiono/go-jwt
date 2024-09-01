package port

import (
	"FaisalBudiono/go-jwt/internal/app/domain"
	"context"
)

type UserRepo interface {
	dbTxMaker

	InsertUser(ctx context.Context, u domain.User, tx DBTx) (domain.User, error)
	FindUserByEmail(ctx context.Context, email string) (domain.User, error)
	FindUserByID(ctx context.Context, userID string) (domain.User, error)
}
