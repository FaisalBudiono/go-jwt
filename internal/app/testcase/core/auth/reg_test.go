package auth_test

import (
	"FaisalBudiono/go-jwt/internal/app/adapter/sqliterepo"
	"FaisalBudiono/go-jwt/internal/app/core"
	"FaisalBudiono/go-jwt/internal/app/core/nower"
	"FaisalBudiono/go-jwt/internal/app/domain"
	"FaisalBudiono/go-jwt/internal/app/testcase"
	"FaisalBudiono/go-jwt/internal/db/sqlc/sqlite/sqlcm"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRegisterUser(t *testing.T) {
	db := testcase.DBConn(testcase.DBMemory)
	now := time.Now()

	err := testcase.NewMigrator(db).Migrate("../../../../db/sqlite_migrations")
	require.Nil(t, err)

	port := &mockRegisterPort{
		ctx:      context.Background(),
		name:     "john doe",
		email:    "johndoe@gmail.com",
		password: "123456",
	}
	deps := newDepsRegister()
	mockHashed, _ := deps.Hash(port.password)

	c := core.NewAuth(
		db,
		deps,
		sqliterepo.New(db, nower.NewFake(time.Now())),
		deps,
		deps,
	)

	got, err := c.Reg(port)

	require.Nil(t, err)
	want := domain.User{
		ID:        got.ID,
		Name:      port.name,
		Email:     port.email,
		Password:  mockHashed,
		CreatedAt: time.Unix(now.Unix(), 0),
		UpdatedAt: time.Unix(now.Unix(), 0),
	}
	assert.Equal(t, want, got)

	dbU, err := sqlcm.New(db).FindUserByEmail(context.Background(), port.email)
	require.Nil(t, err)
	assert.Equal(t, port.name, dbU.Name)
	assert.Equal(t, port.email, dbU.Email)
	assert.Equal(t, mockHashed, dbU.Password)
}

type mockRegisterPort struct {
	ctx      context.Context
	name     string
	email    string
	password string
}

func (m *mockRegisterPort) Ctx() (context.Context, error) {
	return m.ctx, nil
}

func (m *mockRegisterPort) Email() (string, error) {
	return m.email, nil
}

func (m *mockRegisterPort) Name() (string, error) {
	return m.name, nil
}

func (m *mockRegisterPort) Password() (string, error) {
	return m.password, nil
}

type depsRegister struct{}

func (m *depsRegister) Cache(ctx context.Context, t domain.Token) error {
	panic("unimplemented")
}

func (m *depsRegister) Gen(u domain.User) (domain.Token, error) {
	panic("unimplemented")
}

func newDepsRegister() *depsRegister {
	return &depsRegister{}
}

func (m *depsRegister) Hash(plain string) (string, error) {
	return plain + "-hash", nil
}

func (m *depsRegister) Verify(plain string, hashed string) (bool, error) {
	panic("unimplemented")
}
