package jwt

import (
	"FaisalBudiono/go-jwt/internal/app/domain"

	"github.com/go-errors/errors"
)

var ErrTokenExpired = errors.New("JWT token expired")

type TokenManager interface {
	Gen(u domain.User) (domain.Token, error)
	Verify(accessToken string) (userID string, err error)
}
