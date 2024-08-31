package core

import (
	"FaisalBudiono/go-jwt/internal/app/core/hasher"
	"FaisalBudiono/go-jwt/internal/app/core/jwt"
	"FaisalBudiono/go-jwt/internal/app/domain"
	"FaisalBudiono/go-jwt/internal/app/port"
	"FaisalBudiono/go-jwt/internal/app/port/in"
	"database/sql"

	"github.com/go-errors/errors"
)

var ErrInvalidCredential = errors.New("invalid credentials")

func NewAuth(
	userRepo port.UserRepo,
	hasher hasher.PwHasher,
	tokenMng jwt.TokenManager,
	tokenCacher jwt.TokenCacher,
) *auth {
	return &auth{
		userRepo:    userRepo,
		hasher:      hasher,
		tokenMng:    tokenMng,
		tokenCacher: tokenCacher,
	}
}

type auth struct {
	userRepo    port.UserRepo
	hasher      hasher.PwHasher
	tokenMng    jwt.TokenManager
	tokenCacher jwt.TokenCacher
}

func (a *auth) Reg(port in.RegisterPort) (domain.User, error) {
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

	tx, err := a.userRepo.BeginTx(ctx)
	if err != nil {
		return domain.User{}, err
	}
	defer tx.Rollback()

	hash, err := a.hasher.Hash(password)
	if err != nil {
		return domain.User{}, err
	}

	u, err := a.userRepo.InsertUser(ctx, domain.User{
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

func (a *auth) Login(port in.LoginPort) (domain.Token, error) {
	ctx, err := port.Ctx()
	if err != nil {
		return domain.Token{}, err
	}
	email, err := port.Email()
	if err != nil {
		return domain.Token{}, err
	}
	password, err := port.Password()
	if err != nil {
		return domain.Token{}, err
	}

	u, err := a.userRepo.FindUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Token{}, errors.Join(ErrInvalidCredential, err)
		}

		return domain.Token{}, err
	}

	ok, err := a.hasher.Verify(password, u.Password)
	if !ok {
		return domain.Token{}, errors.Join(ErrInvalidCredential, err)
	}

	token, err := a.tokenMng.Gen(u)
	if err != nil {
		return domain.Token{}, err
	}

	err = a.tokenCacher.Cache(ctx, token)
	if err != nil {
		return domain.Token{}, err
	}

	return token, nil
}
