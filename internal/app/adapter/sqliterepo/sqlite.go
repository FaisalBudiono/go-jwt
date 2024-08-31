package sqliterepo

import (
	"FaisalBudiono/go-jwt/internal/app/core/nower"
	"FaisalBudiono/go-jwt/internal/app/domain"
	"FaisalBudiono/go-jwt/internal/db/sqlc/sqlite/sqlcm"
	"context"
	"database/sql"
	"strconv"
	"time"
)

type sqlite struct {
	db    *sql.DB
	nower nower.Nower
}

func New(
	db *sql.DB,
	nower nower.Nower,
) *sqlite {
	return &sqlite{
		db:    db,
		nower: nower,
	}
}

func (s *sqlite) FindUserByEmail(ctx context.Context, email string) (domain.User, error) {
	res, err := s.makeQuery(nil).FindUserByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}

	return domain.User{
		ID:        strconv.FormatInt(res.ID, 10),
		Name:      res.Name,
		Email:     res.Email,
		Password:  res.Password,
		CreatedAt: time.Unix(res.CreatedAt.Int64, 0),
		UpdatedAt: time.Unix(res.UpdatedAt.Int64, 0),
	}, nil
}

func (s *sqlite) InsertUser(ctx context.Context, u domain.User, tx *sql.Tx) (domain.User, error) {
	res, err := s.makeQuery(tx).InsertUser(ctx, sqlcm.InsertUserParams{
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
		CreatedAt: sql.NullInt64{
			Int64: s.nower.Now().Unix(),
			Valid: true,
		},
		UpdatedAt: sql.NullInt64{
			Int64: s.nower.Now().Unix(),
			Valid: true,
		},
	})
	if err != nil {
		return domain.User{}, err
	}

	return domain.User{
		ID:        strconv.FormatInt(res.ID, 10),
		Name:      res.Name,
		Email:     res.Email,
		Password:  res.Password,
		CreatedAt: time.Unix(res.CreatedAt.Int64, 0),
		UpdatedAt: time.Unix(res.UpdatedAt.Int64, 0),
	}, nil
}

func (s *sqlite) makeQuery(tx *sql.Tx) *sqlcm.Queries {
	if tx == nil {
		return sqlcm.New(s.db)
	}

	return sqlcm.New(s.db).WithTx(tx)
}
