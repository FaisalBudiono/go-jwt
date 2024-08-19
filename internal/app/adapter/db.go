package adapter

import (
	"FaisalBudiono/go-jwt/database/sqlc/pg/sqlcm"
	"FaisalBudiono/go-jwt/internal/app/port"
	"FaisalBudiono/go-jwt/internal/app/port/common"
	"context"
	"database/sql"
	"strconv"
)

func NewDB(db *sql.DB) port.DB {
	return &DBSqlc{db: db}
}

type DBSqlc struct {
	db *sql.DB
}

func (d *DBSqlc) DB() *sql.DB {
	return d.db
}

func (d *DBSqlc) FindUserByEmail(ctx context.Context, email string, tx *sql.Tx) (common.User, error) {
	u, err := d.query(tx).FindUserByEmail(ctx, email)
	if err != nil {
		if err != sql.ErrNoRows {
			return common.User{}, err
		}
		return common.User{}, port.ErrDBResNotFound
	}

	return common.User{
		ID:        strconv.FormatInt(u.ID, 10),
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password,
		CreatedAt: u.CreatedAt.Time,
		UpdatedAt: u.UpdatedAt.Time,
	}, nil
}

func (d *DBSqlc) InsertUser(ctx context.Context, u common.User, tx *sql.Tx) (common.User, error) {
	udb, err := d.query(tx).InsertUser(ctx, sqlcm.InsertUserParams{
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	})
	if err != nil {
		return common.User{}, err
	}

	return common.User{
		ID:        strconv.FormatInt(udb.ID, 10),
		Name:      udb.Name,
		Email:     udb.Email,
		Password:  udb.Password,
		CreatedAt: udb.CreatedAt.Time,
		UpdatedAt: udb.UpdatedAt.Time,
	}, nil
}

func (d *DBSqlc) query(tx *sql.Tx) *sqlcm.Queries {
	return sqlcm.New(d.db).WithTx(tx)
}
