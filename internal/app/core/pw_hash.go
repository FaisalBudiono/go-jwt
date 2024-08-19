package core

import "FaisalBudiono/go-jwt/internal/app/core/pwhash/argon"

type PwHasher interface {
	Hash(plain string) (string, error)
	Verify(plain, hashed string) (bool, error)
}

func NewPwHasher() PwHasher {
	return argon.New()
}
