package core_test

import (
	"FaisalBudiono/go-jwt/internal/app/core"
	"FaisalBudiono/go-jwt/internal/app/core/jwt"
	"FaisalBudiono/go-jwt/internal/app/domain"
	mjwt "FaisalBudiono/go-jwt/mocks/internal_/app/core/jwt"
	mport "FaisalBudiono/go-jwt/mocks/internal_/app/port"
	mportin "FaisalBudiono/go-jwt/mocks/internal_/app/port/in"
	"testing"

	"github.com/go-errors/errors"
	"github.com/stretchr/testify/suite"
)

type TokenTestSuite struct {
	suite.Suite

	tokenManager *mjwt.TokenManager
	userRepo     *mport.UserRepo
}

func TestTokenTestSuite(t *testing.T) {
	suite.Run(t, new(TokenTestSuite))
}

func (s *TokenTestSuite) SetupSubTest() {
	s.tokenManager = mjwt.NewTokenManager(s.T())
	s.userRepo = mport.NewUserRepo(s.T())
}

func (s *TokenTestSuite) TestVerify() {
	s.Run("should return error when token is invalid", func() {
		accessToken := "accessToken"

		mockedErr := errors.New("mocked error")

		port := mportin.NewVerifyTokenPort(s.T())
		port.EXPECT().AccessToken().Return(accessToken, nil)

		s.tokenManager.EXPECT().Verify(accessToken).Return("", mockedErr)

		c := core.NewToken(s.tokenManager)
		got, err := c.Verify(port)

		s.Equal(domain.User{}, got)
		s.Equal(mockedErr, err)
	})

	s.Run("should return error when token is expired", func() {
		accessToken := "accessToken"

		port := mportin.NewVerifyTokenPort(s.T())
		port.EXPECT().AccessToken().Return(accessToken, nil)

		s.tokenManager.EXPECT().Verify(accessToken).Return("", jwt.ErrTokenExpired)

		c := core.NewToken(s.tokenManager)
		got, err := c.Verify(port)

		s.Equal(domain.User{}, got)
		s.NotEqual(core.ErrTokenExpired, err)
		s.True(errors.Is(err, core.ErrTokenExpired))
	})

	s.Run("should return error when user not found", func() {
		accessToken := "accessToken"

		mockedUserID := "123"

		port := mportin.NewVerifyTokenPort(s.T())
		port.EXPECT().AccessToken().Return(accessToken, nil)

		s.tokenManager.EXPECT().Verify(accessToken).Return(mockedUserID, nil)

		c := core.NewToken(s.tokenManager)
		got, err := c.Verify(port)

		s.Equal(domain.User{}, got)
		s.NoError(err)
	})

	s.Run("should return user when token is still valid", func() {
		s.T().Skip()
		accessToken := "accessToken"

		port := mportin.NewVerifyTokenPort(s.T())
		port.EXPECT().AccessToken().Return(accessToken, nil)

		s.tokenManager.EXPECT().Verify(accessToken).Return("123", nil)

		c := core.NewToken(s.tokenManager)
		got, err := c.Verify(port)

		s.Equal(domain.User{}, got)
		s.NoError(err)
	})
}
