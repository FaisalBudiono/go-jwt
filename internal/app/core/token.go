package core

import (
	"FaisalBudiono/go-jwt/internal/app/core/jwt"
	"FaisalBudiono/go-jwt/internal/app/domain"
	"FaisalBudiono/go-jwt/internal/app/port/in"

	"github.com/go-errors/errors"
)

var ErrTokenExpired = errors.New("JWT token expired")

type token struct {
	jwtManager jwt.TokenManager
}

func NewToken(
	jwtManager jwt.TokenManager,
) *token {
	return &token{
		jwtManager: jwtManager,
	}
}

func (t *token) Verify(port in.VerifyTokenPort) (domain.User, error) {
	accessToken, err := port.AccessToken()
	if err != nil {
		return domain.User{}, err
	}

	_, err = t.jwtManager.Verify(accessToken)
	if err != nil {
		return domain.User{}, t.mapExpiredErr(err)
	}

	return domain.User{}, nil
}

func (t *token) mapExpiredErr(err error) error {
	if !errors.Is(err, jwt.ErrTokenExpired) {
		return err
	}
	return errors.Join(err, ErrTokenExpired)
}
