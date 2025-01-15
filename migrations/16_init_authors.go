package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table authors")
		_, err := db.Exec(`
			CREATE TABLE IF NOT EXISTS authors (
				id                     SERIAL PRIMARY KEY,
				title                  varchar(255),
				description            jsonb,
			    name                   varchar(255) NOT NULL,
			    slug                   varchar(255),
			    image_id               INTEGER,
				enabled                bool NOT NULL DEFAULT TRUE,
			    order_by               integer,
			    created_by_id          integer,
				updated_by_id          integer,
				created                TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
				updated                TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
				deleted                TIMESTAMP WITH TIME ZONE DEFAULT NULL
			);
			ALTER TABLE authors ADD CONSTRAINT authors_image_id_foreign
				FOREIGN KEY (image_id) REFERENCES files(id);
		`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping authors")
		_, err := db.Exec(`DROP TABLE IF EXISTS authors`)
		return err
	})
}
