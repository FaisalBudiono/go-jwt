package cmd

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

const (
	prodDir   string = "internal/db/migrations"
	sqliteDir string = "internal/db/sqlite_migrations"
)

type gooseMigrator struct {
	db *sql.DB
}

func NewMigrator(db *sql.DB) *gooseMigrator {
	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	return &gooseMigrator{
		db: db,
	}
}

func (g *gooseMigrator) Create(filename string) {
	err := goose.Create(g.db, prodDir, filename, "sql")
	if err != nil {
		panic(err)
	}

	err = goose.Create(g.db, sqliteDir, filename, "sql")
	if err != nil {
		panic(err)
	}
}

func (g *gooseMigrator) Status() {
	err := goose.Status(g.db, prodDir)
	if err != nil {
		panic(err)
	}
}

func (g *gooseMigrator) Version() {
	err := goose.Version(g.db, prodDir)
	if err != nil {
		panic(err)
	}
}

func (g *gooseMigrator) Up() {
	err := goose.Up(g.db, prodDir)
	if err != nil {
		panic(err)
	}
}

func (g *gooseMigrator) Down() {
	err := goose.Down(g.db, prodDir)
	if err != nil {
		panic(err)
	}
}
