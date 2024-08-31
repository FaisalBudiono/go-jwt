package jwt

import (
	"FaisalBudiono/go-jwt/internal/app/domain"
)

type TokenManager interface {
	Gen(u domain.User) (domain.Token, error)
}
