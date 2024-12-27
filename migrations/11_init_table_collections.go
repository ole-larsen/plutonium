package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table collections")

		_, err := db.Exec(`
CREATE TABLE IF NOT EXISTS collections (
	id                     INTEGER PRIMARY KEY,
	name                   varchar(255) UNIQUE,
	slug                   varchar(255),
	url                    varchar(255),
	symbol                 varchar(128),
	description            text,
	fee                    varchar(255),
	is_approved            bool,
	is_locked              bool,
	address                varchar(255),
	metadata               jsonb,
	category_id            INTEGER,
	max_items              INTEGER,
	owner_id               INTEGER,
	creator_id             INTEGER,
	logo_id                INTEGER,
	featured_id            INTEGER,
	banner_id              INTEGER,
	created                TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
	updated                TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
	deleted                TIMESTAMP WITH TIME ZONE DEFAULT NULL
);
ALTER TABLE collections ADD CONSTRAINT collections_category_id_foreign
	FOREIGN KEY (category_id) REFERENCES categories(id);
ALTER TABLE collections ADD CONSTRAINT collections_owner_id_foreign
	FOREIGN KEY (owner_id) REFERENCES users(id);
ALTER TABLE collections ADD CONSTRAINT collections_creator_id_foreign
	FOREIGN KEY (creator_id) REFERENCES users(id);
ALTER TABLE collections ADD CONSTRAINT collections_logo_id_foreign
	FOREIGN KEY (logo_id) REFERENCES files(id);
ALTER TABLE collections ADD CONSTRAINT collections_featured_id_foreign
	FOREIGN KEY (featured_id) REFERENCES files(id);
ALTER TABLE collections ADD CONSTRAINT collections_banner_id_foreign
	FOREIGN KEY (banner_id) REFERENCES files(id);
		`)

		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping collections")

		_, err := db.Exec(`DROP TABLE IF EXISTS collections`)

		return err
	})
}
