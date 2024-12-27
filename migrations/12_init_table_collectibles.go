package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table collectibles")

		_, err := db.Exec(`
CREATE TABLE IF NOT EXISTS collectibles (
	id                     SERIAL PRIMARY KEY,
	item_id                INTEGER,
	token_ids              INTEGER[],
	collection_id          INTEGER,
	uri                    varchar(255),
	creator_id             INTEGER,
	owner_id               INTEGER,
	metadata               jsonb,
	details                jsonb,
	created                TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
	updated                TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
	deleted                TIMESTAMP WITH TIME ZONE DEFAULT NULL
);

ALTER TABLE collectibles ADD CONSTRAINT collectibles_collection_id_foreign
	FOREIGN KEY (collection_id) REFERENCES collections(id);
ALTER TABLE collectibles ADD CONSTRAINT collectibles_owner_id_foreign
	FOREIGN KEY (owner_id) REFERENCES users(id);
ALTER TABLE collectibles ADD CONSTRAINT collectibles_creator_id_foreign
	FOREIGN KEY (creator_id) REFERENCES users(id);
		`)

		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping collectibles")

		_, err := db.Exec(`DROP TABLE IF EXISTS collectibles`)

		return err
	})
}
