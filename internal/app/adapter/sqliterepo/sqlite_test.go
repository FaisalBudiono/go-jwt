package sqliterepo_test

import (
	"FaisalBudiono/go-jwt/internal/app/adapter/sqliterepo"
	"FaisalBudiono/go-jwt/internal/app/domain"
	"FaisalBudiono/go-jwt/internal/app/testcase"
	"FaisalBudiono/go-jwt/internal/db/sqlc/sqlite/sqlcm"
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

const migrationDir = "../../../db/sqlite_migrations"

type SQLiteSuite struct {
	suite.Suite

	db *sql.DB
}

func TestSQLiteSuite(t *testing.T) {
	suite.Run(t, new(SQLiteSuite))
}

func (s *SQLiteSuite) SetupTest() {
	s.db = testcase.DBConn(testcase.DBMemory)

	err := testcase.NewMigrator(s.db).Migrate(migrationDir)
	s.Require().Nil(err)
}

func (s *SQLiteSuite) TearDownSubTest() {
	sqlcm.New(s.db).TruncateUsers(context.Background())
}

func (s *SQLiteSuite) TestInsertUser() {
	s.Run("should successfully rollback insert user", func() {
		u := domain.User{
			Name:     "john doe",
			Email:    "john@doe.com",
			Password: "123456",
		}
		now := time.Now()

		tx, err := s.db.BeginTx(context.Background(), nil)
		s.Require().Nil(err)

		q := sqliterepo.New(s.db, sqliterepo.NewNowerSpy(now))
		_, err = q.InsertUser(context.Background(), u, tx)
		s.Require().Nil(err)

		err = tx.Rollback()
		s.Require().Nil(err)

		dbUsers, err := sqlcm.New(s.db).AllUsers(context.Background())
		s.Require().Nil(err)

		s.Assert().Equal(0, len(dbUsers))
	})

	s.Run("should successfully commit insert user", func() {
		u := domain.User{
			Name:     "john doe",
			Email:    "johndoe@gmail.com",
			Password: "123456",
		}
		now := time.Now()

		tx, err := s.db.BeginTx(context.Background(), nil)
		s.Require().Nil(err)

		q := sqliterepo.New(s.db, sqliterepo.NewNowerSpy(now))
		res, err := q.InsertUser(context.Background(), u, tx)
		s.Require().Nil(err)

		err = tx.Commit()
		s.Require().Nil(err)

		dbUsers, err := sqlcm.New(s.db).AllUsers(context.Background())
		s.Require().Nil(err)

		s.Assert().Equal(1, len(dbUsers))
		s.Assert().Equal(domain.User{
			ID:        res.ID,
			Name:      u.Name,
			Email:     u.Email,
			Password:  u.Password,
			CreatedAt: time.Unix(now.Unix(), 0),
			UpdatedAt: time.Unix(now.Unix(), 0),
		}, res)
	})

	s.Run("should successfully insert user without tx", func() {
		u := domain.User{
			Name:     "john doe",
			Email:    "johndoe@gmail.com",
			Password: "123456",
		}
		now := time.Now()

		q := sqliterepo.New(s.db, sqliterepo.NewNowerSpy(now))
		res, err := q.InsertUser(context.Background(), u, nil)

		s.Require().Nil(err)

		dbUsers, err := sqlcm.New(s.db).AllUsers(context.Background())
		s.Require().Nil(err)

		s.Assert().Equal(1, len(dbUsers))
		s.Assert().Equal(domain.User{
			ID:        res.ID,
			Name:      u.Name,
			Email:     u.Email,
			Password:  u.Password,
			CreatedAt: time.Unix(now.Unix(), 0),
			UpdatedAt: time.Unix(now.Unix(), 0),
		}, res)
	})
}
