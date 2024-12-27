package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table contracts")

		_, err := db.Exec(`
CREATE TABLE IF NOT EXISTS contracts (
	id                     SERIAL PRIMARY KEY,
	name                   varchar(255) UNIQUE NOT NULL,
	address                varchar(255) UNIQUE NOT NULL,
	tx                     varchar(255) UNIQUE NOT NULL,
	abi                    text,
	created                TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
	updated                TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
	deleted                TIMESTAMP WITH TIME ZONE DEFAULT NULL
);
		`)

		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping contracts")

		_, err := db.Exec(`DROP TABLE IF EXISTS contracts`)

		return err
	})
}
