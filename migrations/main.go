package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ole-larsen/plutonium/internal/plutonium/settings"

	"github.com/go-pg/migrations"
	"github.com/go-pg/pg"
)

const usageText = `This program runs command on the db. Supported commands are:
  - up - runs all available migrations.
  - up [target] - runs available migrations up to the target one.
  - down - reverts last migration.
  - reset - reverts all migrations.
  - version - prints current db version.
  - set_version [version] - sets db version without running migrations.

Usage:
  go run *.go <command> [args]
`

func main() {
	flag.Usage = usage
	flag.Parse()

	cfg := settings.LoadConfig("../.env")

	if cfg == nil {
		exitf("config not found")
	}

	db := pg.Connect(&pg.Options{
		Addr:     cfg.DB.Host + ":" + cfg.DB.Port,
		User:     cfg.DB.Username,
		Password: cfg.DB.Password,
		Database: cfg.DB.Database,
	})

	oldVersion, newVersion, err := migrations.Run(db, flag.Args()...)
	if err != nil {
		exitf(err.Error())
	}

	if newVersion != oldVersion {
		fmt.Printf("migrated from version %d to %d\n", oldVersion, newVersion)
	} else {
		fmt.Printf("version is %d\n", oldVersion)
	}
}

func usage() {
	exitCode := 2

	fmt.Print(usageText)
	flag.PrintDefaults()
	os.Exit(exitCode)
}

func errorf(s string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, s+"\n", args...)
}

func exitf(s string, args ...interface{}) {
	errorf(s, args...)
	os.Exit(1)
}
