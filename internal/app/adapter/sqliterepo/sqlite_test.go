package sqliterepo_test

import (
	"FaisalBudiono/go-jwt/internal/app/adapter/sqliterepo"
	"FaisalBudiono/go-jwt/internal/app/domain"
	"FaisalBudiono/go-jwt/internal/app/testcase"
	"FaisalBudiono/go-jwt/internal/db/sqlc/sqlite/sqlcm"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const migrationDir = "../../../db/sqlite_migrations"

func TestInsertUser(t *testing.T) {
	t.Run("should successfully rollback insert user", func(t *testing.T) {
		db := testcase.DBConn(testcase.DBMemory)

		err := testcase.NewMigrator(db).Migrate(migrationDir)
		require.Nil(t, err)

		u := domain.User{
			Name:     "john doe",
			Email:    "john@doe.com",
			Password: "123456",
		}
		now := time.Now()

		tx, err := db.BeginTx(context.Background(), nil)
		require.Nil(t, err)

		q := sqliterepo.New(db, newNowerMock(now))
		_, err = q.InsertUser(context.Background(), u, tx)
		require.Nil(t, err)

		err = tx.Rollback()
		require.Nil(t, err)

		dbUsers, err := sqlcm.New(db).AllUsers(context.Background())
		require.Nil(t, err)

		assert.Equal(t, 0, len(dbUsers))
	})

	t.Run("should successfully commit insert user", func(t *testing.T) {
		db := testcase.DBConn(testcase.DBMemory)

		err := testcase.NewMigrator(db).Migrate(migrationDir)
		require.Nil(t, err)

		u := domain.User{
			Name:     "john doe",
			Email:    "johndoe@gmail.com",
			Password: "123456",
		}
		now := time.Now()

		tx, err := db.BeginTx(context.Background(), nil)
		require.Nil(t, err)

		q := sqliterepo.New(db, newNowerMock(now))
		res, err := q.InsertUser(context.Background(), u, tx)
		require.Nil(t, err)

		err = tx.Commit()
		require.Nil(t, err)

		dbUsers, err := sqlcm.New(db).AllUsers(context.Background())
		require.Nil(t, err)

		assert.Equal(t, 1, len(dbUsers))
		assert.Equal(t, domain.User{
			ID:        res.ID,
			Name:      u.Name,
			Email:     u.Email,
			Password:  u.Password,
			CreatedAt: time.Unix(now.Unix(), 0),
			UpdatedAt: time.Unix(now.Unix(), 0),
		}, res)
	})

	t.Run("should successfully insert user without tx", func(t *testing.T) {
		db := testcase.DBConn(testcase.DBMemory)

		err := testcase.NewMigrator(db).Migrate(migrationDir)
		require.Nil(t, err)

		u := domain.User{
			Name:     "john doe",
			Email:    "johndoe@gmail.com",
			Password: "123456",
		}
		now := time.Now()

		q := sqliterepo.New(db, newNowerMock(now))
		res, err := q.InsertUser(context.Background(), u, nil)

		require.Nil(t, err)

		dbUsers, err := sqlcm.New(db).AllUsers(context.Background())
		require.Nil(t, err)

		assert.Equal(t, 1, len(dbUsers))
		assert.Equal(t, domain.User{
			ID:        res.ID,
			Name:      u.Name,
			Email:     u.Email,
			Password:  u.Password,
			CreatedAt: time.Unix(now.Unix(), 0),
			UpdatedAt: time.Unix(now.Unix(), 0),
		}, res)
	})
}

type nowerMock struct {
	now time.Time
}

func newNowerMock(t time.Time) *nowerMock {
	return &nowerMock{
		now: t,
	}
}

func (n *nowerMock) Now() time.Time {
	return n.now
}
