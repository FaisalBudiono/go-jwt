package testcase

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"
)

const (
	DBMemory string = ":memory:"
)

func DBConn(conn string) *sql.DB {
	db, err := sql.Open("sqlite3", conn)
	if err != nil {
		panic(err)
	}

	return db
}

type migrator struct {
	db *sql.DB
}

func NewMigrator(db *sql.DB) *migrator {
	return &migrator{
		db: db,
	}
}

func (m *migrator) Migrate(migrationDir string) error {
	err := goose.SetDialect("sqlite3")
	if err != nil {
		return err
	}

	err = goose.Up(m.db, migrationDir)
	if err != nil {
		return fmt.Errorf("failed migration .. %w", err)
	}

	return nil
}
