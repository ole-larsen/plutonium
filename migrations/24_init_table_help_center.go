package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table help_center")

		_, err := db.Exec(`
CREATE TABLE IF NOT EXISTS help_center (
	id                     SERIAL PRIMARY KEY,
	title                  varchar(255) UNIQUE NOT NULL,
	slug                   varchar(255),
	description            text,
	image_id               INTEGER NOT NULL,
	enabled                bool NOT NULL DEFAULT TRUE,
	order_by               integer,
	created_by_id          integer,
	updated_by_id          integer,
	created                TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
	updated                TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
	deleted                TIMESTAMP WITH TIME ZONE DEFAULT NULL
);
ALTER TABLE help_center ADD CONSTRAINT help_center_image_id_foreign
	FOREIGN KEY (image_id) REFERENCES files(id);
ALTER TABLE help_center ADD CONSTRAINT help_center_created_by_id_foreign
	FOREIGN KEY (created_by_id) REFERENCES users(id);
ALTER TABLE help_center ADD CONSTRAINT help_center_updated_by_id_foreign
	FOREIGN KEY (updated_by_id) REFERENCES users(id);
		`)

		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping help_center")

		_, err := db.Exec(`DROP TABLE IF EXISTS help_center`)

		return err
	})
}
