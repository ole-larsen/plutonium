package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table contacts")

		_, err := db.Exec(`
CREATE TABLE IF NOT EXISTS contacts (
	id                     SERIAL PRIMARY KEY,
	page_id                INTEGER,
	heading                text,
	sub_heading            text,
	image_id               INTEGER,
	enabled                bool NOT NULL DEFAULT TRUE,
	created_by_id          integer,
	updated_by_id          integer,
	created                TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
	updated                TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
	deleted                TIMESTAMP WITH TIME ZONE DEFAULT NULL
);
ALTER TABLE contacts ADD CONSTRAINT contacts_page_id_foreign
	FOREIGN KEY (page_id) REFERENCES pages(id);
ALTER TABLE contacts ADD CONSTRAINT contacts_image_id_foreign
	FOREIGN KEY (image_id) REFERENCES files(id);
ALTER TABLE contacts ADD CONSTRAINT contacts_created_by_id_foreign
	FOREIGN KEY (created_by_id) REFERENCES users(id);
ALTER TABLE contacts ADD CONSTRAINT contacts_updated_by_id_foreign
	FOREIGN KEY (updated_by_id) REFERENCES users(id);`)

		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping contacts")

		_, err := db.Exec(`DROP TABLE IF EXISTS contacts`)

		return err
	})
}
