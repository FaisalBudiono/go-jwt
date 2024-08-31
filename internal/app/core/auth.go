package core

import (
	"FaisalBudiono/go-jwt/internal/app/core/hasher"
	"FaisalBudiono/go-jwt/internal/app/domain"
	"context"
	"database/sql"

	"github.com/go-errors/errors"
)

var ErrInvalidCredential = errors.New("invalid credentials")

func NewAuth(
	db *sql.DB,
	hasher hasher.PwHasher,
	uRepo userRepo,
	tokenGen tokenGenerator,
	tokenC tokenCacher,
) *auth {
	return &auth{
		db:       db,
		hasher:   hasher,
		uRepo:    uRepo,
		tokenGen: tokenGen,
		tokenC:   tokenC,
	}
}

type (
	userRepo interface {
		InsertUser(ctx context.Context, u domain.User, tx *sql.Tx) (domain.User, error)
		FindUserByEmail(ctx context.Context, email string) (domain.User, error)
	}
	tokenGenerator interface {
		Gen(u domain.User) (domain.Token, error)
	}
	tokenCacher interface {
		Cache(ctx context.Context, t domain.Token) error
	}
)

type auth struct {
	db       *sql.DB
	hasher   hasher.PwHasher
	uRepo    userRepo
	tokenGen tokenGenerator
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

	u, err := a.uRepo.InsertUser(ctx, domain.User{
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

	u, err := a.uRepo.FindUserByEmail(ctx, email)
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

	token, err := a.tokenGen.Gen(u)
	if err != nil {
		return domain.Token{}, err
	}

	err = a.tokenC.Cache(ctx, token)
	if err != nil {
		return domain.Token{}, err
	}

	return token, nil
}
