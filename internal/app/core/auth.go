package core

import (
	"FaisalBudiono/go-jwt/internal/app/core/hasher"
	"FaisalBudiono/go-jwt/internal/app/domain"
	"context"
	"database/sql"
)

func NewAuth(
	db *sql.DB,
	hasher hasher.PwHasher,
	uInserter userInserter,
) *auth {
	return &auth{
		db:        db,
		hasher:    hasher,
		uInserter: uInserter,
	}
}

type userInserter interface {
	InsertUser(ctx context.Context, u domain.User, tx *sql.Tx) (domain.User, error)
}

type auth struct {
	db        *sql.DB
	hasher    hasher.PwHasher
	uInserter userInserter
}

type registerPort interface {
	Ctx() (context.Context, error)
	Name() (string, error)
	Email() (string, error)
	Password() (string, error)
}

func (a *auth) Reg(port registerPort) (domain.User, error) {
	ctx, err := port.Ctx()
	if err != nil {
		return domain.User{}, err
	}
	name, err := port.Name()
	if err != nil {
		return domain.User{}, err
	}
	email, err := port.Email()
	if err != nil {
		return domain.User{}, err
	}
	password, err := port.Password()
	if err != nil {
		return domain.User{}, err
	}

	tx, err := a.db.BeginTx(ctx, nil)
	if err != nil {
		return domain.User{}, err
	}
	defer tx.Rollback()

	hash, err := a.hasher.Hash(password)
	if err != nil {
		return domain.User{}, err
	}

	u, err := a.uInserter.InsertUser(ctx, domain.User{
		Name:     name,
		Email:    email,
		Password: hash,
	}, tx)
	if err != nil {
		return domain.User{}, err
	}

	err = tx.Commit()
	if err != nil {
		return domain.User{}, err
	}

	return u, nil
}
