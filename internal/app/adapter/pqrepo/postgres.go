package pqrepo

import (
	"FaisalBudiono/go-jwt/internal/app/domain"
	"FaisalBudiono/go-jwt/internal/app/port"
	"FaisalBudiono/go-jwt/internal/db/sqlc/pg/sqlcm"
	"context"
	"database/sql"
	"strconv"
)

type postgres struct {
	db *sql.DB
}

func New(
	db *sql.DB,
) *postgres {
	return &postgres{
		db: db,
	}
}

func (p *postgres) BeginTx(ctx context.Context) (port.DBTx, error) {
	tx, err := p.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	return newTx(tx), nil
}

func (p *postgres) InsertUser(ctx context.Context, u domain.User, tx port.DBTx) (domain.User, error) {
	res, err := p.makeQuery(tx).InsertUser(ctx, sqlcm.InsertUserParams{
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	})
	if err != nil {
		return domain.User{}, err
	}

	return domain.User{
		ID:        strconv.FormatInt(res.ID, 10),
		Name:      res.Name,
		Email:     res.Email,
		Password:  res.Password,
		CreatedAt: res.CreatedAt.Time,
		UpdatedAt: res.UpdatedAt.Time,
	}, nil
}

func (p *postgres) FindUserByEmail(ctx context.Context, email string) (domain.User, error) {
	res, err := p.makeQuery(nil).FindUserByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}

	return domain.User{
		ID:        strconv.FormatInt(res.ID, 10),
		Name:      res.Name,
		Email:     res.Email,
		Password:  res.Password,
		CreatedAt: res.CreatedAt.Time,
		UpdatedAt: res.UpdatedAt.Time,
	}, nil
}

func (p *postgres) makeQuery(tx port.DBTx) *sqlcm.Queries {
	if tx == nil {
		return sqlcm.New(p.db)
	}

	sqlTx, ok := tx.(*sqlTx)
	if !ok {
		panic("tx is not sqlTx")
	}

	return sqlcm.New(p.db).WithTx(sqlTx.tx)
}
