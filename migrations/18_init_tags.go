package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table tags")
		_, err := db.Exec(`
			CREATE TABLE IF NOT EXISTS tags (
				id                     SERIAL PRIMARY KEY,
				parent_id              INTEGER NOT NULL,
				title                  varchar(128) UNIQUE NOT NULL,
			    slug                   varchar(128) UNIQUE NOT NULL,
				enabled                bool NOT NULL DEFAULT TRUE,
			    created_by_id          integer,
				updated_by_id          integer,
				created                TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
				updated                TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
				deleted                TIMESTAMP WITH TIME ZONE DEFAULT NULL
			);
		`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping tags")
		_, err := db.Exec(`DROP TABLE IF EXISTS tags`)
		return err
	})
}
