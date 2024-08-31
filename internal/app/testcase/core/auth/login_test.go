package auth_test

import (
	"FaisalBudiono/go-jwt/internal/app/adapter/sqliterepo"
	"FaisalBudiono/go-jwt/internal/app/core"
	"FaisalBudiono/go-jwt/internal/app/core/nower"
	"FaisalBudiono/go-jwt/internal/app/domain"
	"FaisalBudiono/go-jwt/internal/app/testcase"
	"FaisalBudiono/go-jwt/internal/db/sqlc/sqlite/sqlcm"
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type LoginSuite struct {
	suite.Suite

	db      *sql.DB
	regUser domain.User
}

func TestLoginSuite(t *testing.T) {
	suite.Run(t, new(LoginSuite))
}

func (s *LoginSuite) SetupTest() {
	s.db = testcase.DBConn(testcase.DBMemory)

	err := testcase.NewMigrator(s.db).Migrate("../../../../db/sqlite_migrations")
	s.Require().Nil(err)

	s.regUser = domain.User{
		ID:        "1",
		Name:      "John Doe",
		Email:     "johndoe@gmail.com",
		Password:  "123456",
		CreatedAt: time.Unix(time.Now().Unix(), 0),
		UpdatedAt: time.Unix(time.Now().Unix(), 0),
	}

	sqlcm.New(s.db).InsertUser(context.Background(), sqlcm.InsertUserParams{
		Name:     s.regUser.Name,
		Email:    s.regUser.Email,
		Password: s.regUser.Password,
		CreatedAt: sql.NullInt64{
			Int64: s.regUser.CreatedAt.Unix(),
			Valid: true,
		},
		UpdatedAt: sql.NullInt64{
			Int64: s.regUser.UpdatedAt.Unix(),
			Valid: true,
		},
	})
}

func (s *LoginSuite) TestUserNotFoundInDatabase() {
	port := &mockLoginPort{
		ctx:      context.Background(),
		email:    "notjohndoe@gmail.com",
		password: "123456",
	}
	deps := newDepsLogin(s, nil)

	c := core.NewAuth(
		s.db,
		deps,
		sqliterepo.New(s.db, nower.NewFake(time.Now())),
		deps,
		deps,
	)

	got, err := c.Login(port)

	s.Assert().Equal(domain.Token{}, got)

	errStacks := []error{
		core.ErrInvalidCredential,
		sql.ErrNoRows,
	}
	for _, e := range errStacks {
		s.Assert().True(s.Assert().ErrorIs(err, e))
	}
}

func (s *LoginSuite) TestUserNotFoundInDatabaseCauseByGenericError() {
	port := &mockLoginPort{
		ctx:      context.Background(),
		email:    "notjohndoe@gmail.com",
		password: "123456",
	}
	deps := newDepsLogin(s, nil)

	anotherDB := testcase.DBConn(testcase.DBMemory)

	c := core.NewAuth(
		s.db,
		deps,
		sqliterepo.New(anotherDB, nower.NewFake(time.Now())),
		deps,
		deps,
	)

	got, err := c.Login(port)

	s.Assert().Equal(domain.Token{}, got)

	s.Assert().NotNil(err)
	s.Assert().NotEqual(core.ErrInvalidCredential, err)
}

func (s *LoginSuite) TestPasswordNotMatch() {
	port := &mockLoginPort{
		ctx:      context.Background(),
		email:    s.regUser.Email,
		password: "123",
	}
	deps := newDepsLogin(s, nil)

	c := core.NewAuth(
		s.db,
		deps,
		sqliterepo.New(s.db, nower.NewFake(time.Now())),
		deps,
		deps,
	)

	got, err := c.Login(port)

	s.Assert().Equal(domain.Token{}, got)
	s.Assert().NotEqual(core.ErrInvalidCredential, err)
	s.Assert().True(s.Assert().ErrorIs(err, core.ErrInvalidCredential))
}

func (s *LoginSuite) TestSuccessfullyLoginWithValidCredentials() {
	port := &mockLoginPort{
		ctx:      context.Background(),
		email:    s.regUser.Email,
		password: s.regUser.Password,
	}

	deps := newDepsLogin(s, nil)
	mockToken, _ := deps.Gen(s.regUser)
	deps.cachedToken = &mockToken

	c := core.NewAuth(
		s.db,
		deps,
		sqliterepo.New(s.db, nower.NewFake(time.Now())),
		deps,
		deps,
	)

	got, err := c.Login(port)

	s.Require().Nil(err)
	s.Assert().Equal(mockToken, got)
}

type mockLoginPort struct {
	ctx      context.Context
	email    string
	password string
}

func (m *mockLoginPort) Ctx() (context.Context, error) {
	return m.ctx, nil
}

func (m *mockLoginPort) Email() (string, error) {
	return m.email, nil
}

func (m *mockLoginPort) Password() (string, error) {
	return m.password, nil
}

type depsLogin struct {
	suite *LoginSuite

	cachedToken *domain.Token
}

func (d *depsLogin) Cache(ctx context.Context, t domain.Token) error {
	d.suite.Assert().Equal(*d.cachedToken, t)

	return nil
}

func (d *depsLogin) Gen(u domain.User) (domain.Token, error) {
	return domain.Token{
		AccessToken: u.Email + u.Name,
		RefreshToken: domain.RefreshToken{
			RootID:       u.CreatedAt.String(),
			ParentID:     u.CreatedAt.String(),
			RefreshToken: u.CreatedAt.String(),
		},
	}, nil
}

func newDepsLogin(
	suite *LoginSuite,
	cachedToken *domain.Token,
) *depsLogin {
	return &depsLogin{
		suite:       suite,
		cachedToken: cachedToken,
	}
}

func (d *depsLogin) Hash(plain string) (string, error) {
	panic("unimplemented")
}

func (d *depsLogin) Verify(plain string, hashed string) (bool, error) {
	return plain == hashed, nil
}
