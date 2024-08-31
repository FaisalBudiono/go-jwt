package jwt

import (
	"FaisalBudiono/go-jwt/internal/app/domain"
	"context"
)

type TokenCacher interface {
	Cache(ctx context.Context, t domain.Token) error
}
