package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table roles")

		_, err := db.Exec(`
CREATE TABLE IF NOT EXISTS roles (
	id          SERIAL PRIMARY KEY,
	title       varchar(255) UNIQUE NOT NULL,
	description text NOT NULL,
	enabled     bool NOT NULL DEFAULT TRUE,
	created     TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
	updated     TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
	deleted     TIMESTAMP WITH TIME ZONE DEFAULT NULL
);`)

		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping table roles")

		_, err := db.Exec(`DROP TABLE IF EXISTS roles`)

		return err
	})
}
