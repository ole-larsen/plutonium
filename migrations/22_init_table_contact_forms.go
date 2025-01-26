package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table contacts_forms")

		_, err := db.Exec(`
CREATE TABLE IF NOT EXISTS contacts_forms (
	id                     SERIAL PRIMARY KEY,
	page_id                integer,
	provider               varchar(255),
	name                   varchar(255),
	email                  varchar(255),
	subject                varchar(255),
	message                text,
	created_by_id          integer,
	updated_by_id          integer,
	created                TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
	updated                TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
	deleted                TIMESTAMP WITH TIME ZONE DEFAULT NULL
);
ALTER TABLE contacts_forms ADD CONSTRAINT contacts_forms_page_id_foreign
	FOREIGN KEY (page_id) REFERENCES pages(id) ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE contacts_forms ADD CONSTRAINT contacts_forms_created_by_id_foreign
	FOREIGN KEY (created_by_id) REFERENCES users(id);
ALTER TABLE contacts_forms ADD CONSTRAINT contacts_forms_updated_by_id_foreign
	FOREIGN KEY (updated_by_id) REFERENCES users(id);

		`)

		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping pages")

		_, err := db.Exec(`DROP TABLE IF EXISTS contacts_forms`)

		return err
	})
}
