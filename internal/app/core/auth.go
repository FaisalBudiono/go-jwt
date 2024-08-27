package core

import (
	"FaisalBudiono/go-jwt/internal/app/port"
	"FaisalBudiono/go-jwt/internal/app/port/common"
)

type pwHasher interface {
	Hash(plain string) (string, error)
	Verify(plain, hashed string) (bool, error)
}

type auth struct {
	db     port.DB
	hasher pwHasher
}

func NewAuth(db port.DB, hasher pwHasher) *auth {
	return &auth{
		db:     db,
		hasher: hasher,
	}
}

func (a *auth) Reg(p port.RegisterInput) (common.User, error) {
	tx, err := a.db.DB().BeginTx(p.Ctx, nil)
	if err != nil {
		return common.User{}, err
	}
	defer tx.Rollback()

	hash, err := a.hasher.Hash(p.Password)
	if err != nil {
		return common.User{}, err
	}

	u, err := a.db.InsertUser(p.Ctx, common.User{
		Name:     p.Name,
		Email:    p.Email,
		Password: hash,
	}, tx)
	if err != nil {
		return common.User{}, err
	}

	err = tx.Commit()
	if err != nil {
		return common.User{}, err
	}

	return u, nil
}

func (a *auth) Login(p port.LoginInput) (common.To, error) {
	tx, err := a.db.DB().BeginTx(p.Ctx, nil)
	if err != nil {
		return common.User{}, err
	}
	defer tx.Rollback()
}
