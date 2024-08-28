package main

import (
	"FaisalBudiono/go-jwt/cmd/migrator/cmd"
	"FaisalBudiono/go-jwt/internal/env"
	"fmt"
	"os"
)

const (
	create string = "create"
)

func init() {
	env.Bind()
}

func main() {
	args := os.Args
	if len(args) == 1 {
		helpScreen()
		os.Exit(1)
	}

	switch args[1] {
	case cmd.CmdCreate:
		cmd.RunCreate()
	case cmd.CmdDown:
		cmd.RunDown()
	case cmd.CmdStatus:
		cmd.RunStatus()
	case cmd.CmdUp:
		cmd.RunUp()
	case cmd.CmdVersion:
		cmd.RunVersion()
	default:
		helpScreen()
	}
}

func helpScreen() {
	fmt.Printf("Should keyin valid command:\n")
	fmt.Printf("    - %s\n", cmd.CmdCreate)
	fmt.Printf("    - %s\n", cmd.CmdDown)
	fmt.Printf("    - %s\n", cmd.CmdStatus)
	fmt.Printf("    - %s\n", cmd.CmdUp)
	fmt.Printf("    - %s\n", cmd.CmdVersion)
}
