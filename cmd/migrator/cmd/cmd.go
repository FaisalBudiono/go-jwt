package cmd

import (
	"FaisalBudiono/go-jwt/internal/db"
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	CmdCreate  string = "create"
	CmdDown    string = "down"
	CmdStatus  string = "status"
	CmdUp      string = "up"
	CmdVersion string = "version"
)

func RunCreate() {
	fmt.Print("Type migration file name: ")

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	input := sc.Text()

	firstArg := strings.Split(input, " ")[0]

	NewMigrator(db.PostgresConn()).Create(firstArg)

	fmt.Println()
	fmt.Println("Migration file successfully created")
}

func RunDown() {
	fmt.Println("Start rolling back migration...")
	NewMigrator(db.PostgresConn()).Down()
	fmt.Println("Finish rolling back migration...")
}

func RunStatus() {
	NewMigrator(db.PostgresConn()).Status()
}

func RunUp() {
	fmt.Println("Start migrating migration...")
	NewMigrator(db.PostgresConn()).Up()
	fmt.Println("Finish migrating migration...")
}

func RunVersion() {
	NewMigrator(db.PostgresConn()).Version()
}
