package core_test

import (
	"FaisalBudiono/go-jwt/internal/app/core"
	"FaisalBudiono/go-jwt/internal/app/domain"
	mhasher "FaisalBudiono/go-jwt/mocks/internal_/app/core/hasher"
	mjwt "FaisalBudiono/go-jwt/mocks/internal_/app/core/jwt"
	mport "FaisalBudiono/go-jwt/mocks/internal_/app/port"
	mportin "FaisalBudiono/go-jwt/mocks/internal_/app/port/in"
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/go-errors/errors"
	"github.com/stretchr/testify/suite"
)

type AuthRegTestSuite struct {
	suite.Suite

	tx           *mport.DBTx
	userRepo     *mport.UserRepo
	hasher       *mhasher.PwHasher
	tokenManager *mjwt.TokenManager
	tokenCacher  *mjwt.TokenCacher
}

func TestAuthRegTestSuite(t *testing.T) {
	suite.Run(t, new(AuthRegTestSuite))
}

func (s *AuthRegTestSuite) SetupSubTest() {
	s.tx = mport.NewDBTx(s.T())
	s.userRepo = mport.NewUserRepo(s.T())
	s.hasher = mhasher.NewPwHasher(s.T())
	s.tokenManager = mjwt.NewTokenManager(s.T())
	s.tokenCacher = mjwt.NewTokenCacher(s.T())
}

func (s *AuthRegTestSuite) TestReg() {
	s.Run("should return user when successfully registered", func() {
		ctx := context.Background()
		name := "john doe"
		email := "johndoe@gmail.com"
		password := "123456"

		mockHashed := password + "-hash"

		mockedSavedUser := domain.User{
			ID:        "1",
			Name:      name,
			Email:     email,
			Password:  mockHashed,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		s.tx.EXPECT().Commit().Return(nil)
		s.tx.EXPECT().Rollback().Return(nil)

		s.userRepo.EXPECT().BeginTx(ctx).Return(s.tx, nil)
		s.userRepo.EXPECT().InsertUser(ctx, domain.User{
			Name:     name,
			Email:    email,
			Password: mockHashed,
		}, s.tx).Return(mockedSavedUser, nil)

		s.hasher.EXPECT().Hash(password).Return(mockHashed, nil)

		port := mportin.NewRegisterPort(s.T())
		port.EXPECT().Ctx().Return(ctx, nil)
		port.EXPECT().Name().Return(name, nil)
		port.EXPECT().Email().Return(email, nil)
		port.EXPECT().Password().Return(password, nil)

		c := core.NewAuth(s.userRepo, s.hasher, s.tokenManager, s.tokenCacher)
		got, err := c.Reg(port)

		s.Equal(mockedSavedUser, got)
		s.NoError(err)
	})
}

func (s *AuthRegTestSuite) TestLogin() {
	s.Run("return error when user not found", func() {
		ctx := context.Background()
		email := "johndoe@gmail.com"
		password := "123456"

		port := mportin.NewLoginPort(s.T())
		port.EXPECT().Ctx().Return(ctx, nil)
		port.EXPECT().Email().Return(email, nil)
		port.EXPECT().Password().Return(password, nil)

		s.userRepo.EXPECT().FindUserByEmail(ctx, email).Return(domain.User{}, sql.ErrNoRows)

		c := core.NewAuth(s.userRepo, s.hasher, s.tokenManager, s.tokenCacher)
		got, err := c.Login(port)

		s.Equal(domain.Token{}, got)

		errStacks := []error{
			core.ErrInvalidCredential,
			sql.ErrNoRows,
		}
		for _, e := range errStacks {
			s.Assert().True(s.Assert().ErrorIs(err, e))
		}
	})

	s.Run("return error when user not found caused by generic error", func() {
		ctx := context.Background()
		email := "johndoe@gmail.com"
		password := "123456"

		mockedErr := errors.New("generic error")

		port := mportin.NewLoginPort(s.T())
		port.EXPECT().Ctx().Return(ctx, nil)
		port.EXPECT().Email().Return(email, nil)
		port.EXPECT().Password().Return(password, nil)

		s.userRepo.EXPECT().FindUserByEmail(ctx, email).Return(domain.User{}, mockedErr)

		c := core.NewAuth(s.userRepo, s.hasher, s.tokenManager, s.tokenCacher)
		got, err := c.Login(port)

		s.Equal(domain.Token{}, got)
		s.Equal(mockedErr, err)
	})

	s.Run("return error when password invalid", func() {
		ctx := context.Background()
		email := "johndoe@gmail.com"
		password := "123456"

		mockedHashed := password + "-hash"

		foundUser := domain.User{
			ID:        "1",
			Name:      "John Doe",
			Email:     email,
			Password:  mockedHashed,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		port := mportin.NewLoginPort(s.T())
		port.EXPECT().Ctx().Return(ctx, nil)
		port.EXPECT().Email().Return(email, nil)
		port.EXPECT().Password().Return(password, nil)

		s.userRepo.EXPECT().FindUserByEmail(ctx, email).Return(foundUser, nil)

		s.hasher.EXPECT().Verify(password, mockedHashed).Return(false, nil)

		c := core.NewAuth(s.userRepo, s.hasher, s.tokenManager, s.tokenCacher)
		got, err := c.Login(port)

		s.Equal(domain.Token{}, got)
		s.Assert().NotEqual(core.ErrInvalidCredential, err)
		s.Assert().True(s.Assert().ErrorIs(err, core.ErrInvalidCredential))
	})

	s.Run("return token when successfully login with valid credentials", func() {
		ctx := context.Background()
		email := "johndoe@gmail.com"
		password := "123456"

		mockedHashed := password + "-hash"

		foundUser := domain.User{
			ID:        "1",
			Name:      "John Doe",
			Email:     email,
			Password:  mockedHashed,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		mockedToken := domain.Token{
			AccessToken: "access-token",
			RefreshToken: domain.RefreshToken{
				RootID:       "root-id",
				ParentID:     "parent-id",
				RefreshToken: "refresh-token",
			},
		}

		port := mportin.NewLoginPort(s.T())
		port.EXPECT().Ctx().Return(ctx, nil)
		port.EXPECT().Email().Return(email, nil)
		port.EXPECT().Password().Return(password, nil)

		s.userRepo.EXPECT().FindUserByEmail(ctx, email).Return(foundUser, nil)

		s.hasher.EXPECT().Verify(password, mockedHashed).Return(true, nil)

		s.tokenManager.EXPECT().Gen(foundUser).Return(mockedToken, nil)
		s.tokenCacher.EXPECT().Cache(ctx, mockedToken).Return(nil)

		c := core.NewAuth(s.userRepo, s.hasher, s.tokenManager, s.tokenCacher)
		got, err := c.Login(port)

		s.Equal(mockedToken, got)
		s.NoError(err)
	})
}
