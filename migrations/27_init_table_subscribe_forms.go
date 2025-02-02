package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table subscribe_forms")

		_, err := db.Exec(`
			CREATE TABLE IF NOT EXISTS subscribe_forms (
				id                     SERIAL PRIMARY KEY,
				provider               varchar(255),
				email                  varchar(255) UNIQUE NOT NULL,
			    created_by_id          integer,
				updated_by_id          integer,
				created                TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
				updated                TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
				deleted                TIMESTAMP WITH TIME ZONE DEFAULT NULL
			);
		`)

		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping subscribe_forms")

		_, err := db.Exec(`DROP TABLE IF EXISTS subscribe_forms`)

		return err
	})
}
