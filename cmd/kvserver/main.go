package main

import (
	"os"

	"github.com/ishii1648/mem-kvsd/cmd/kvserver/app"
)

func main() {
	command := app.NewKVServerCommand()

	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
