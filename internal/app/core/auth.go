package core

import (
	"FaisalBudiono/go-jwt/internal/app/core/hasher"
	"FaisalBudiono/go-jwt/internal/app/core/jwt"
	"FaisalBudiono/go-jwt/internal/app/domain"
	"FaisalBudiono/go-jwt/internal/app/port"
	"context"
	"database/sql"

	"github.com/go-errors/errors"
)

var ErrInvalidCredential = errors.New("invalid credentials")

func NewAuth(
	db *sql.DB,
	hasher hasher.PwHasher,
	userRepo port.UserRepo,
	tokenMng jwt.TokenManager,
	tokenC tokenCacher,
) *auth {
	return &auth{
		db:       db,
		hasher:   hasher,
		userRepo: userRepo,
		tokenMng: tokenMng,
		tokenC:   tokenC,
	}
}

type (
	tokenCacher interface {
		Cache(ctx context.Context, t domain.Token) error
	}
)

type auth struct {
	db       *sql.DB
	hasher   hasher.PwHasher
	userRepo port.UserRepo
	tokenMng jwt.TokenManager
	tokenC   tokenCacher
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

type loginPort interface {
	Ctx() (context.Context, error)
	Email() (string, error)
	Password() (string, error)
}

func (a *auth) Login(port loginPort) (domain.Token, error) {
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

	err = a.tokenC.Cache(ctx, token)
	if err != nil {
		return domain.Token{}, err
	}

	return token, nil
}
