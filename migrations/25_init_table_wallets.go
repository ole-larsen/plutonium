package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table wallets")

		_, err := db.Exec(`
CREATE TABLE IF NOT EXISTS wallets (
	id                     SERIAL PRIMARY KEY,
	title                  varchar(128) UNIQUE NOT NULL,
	description            text,
	image_id               INTEGER,
	enabled                bool NOT NULL DEFAULT TRUE,
	order_by               integer,
	created_by_id          integer,
	updated_by_id          integer,
	created                TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
	updated                TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
	deleted                TIMESTAMP WITH TIME ZONE DEFAULT NULL
);
ALTER TABLE wallets ADD CONSTRAINT wallets_image_id_foreign
	FOREIGN KEY (image_id) REFERENCES files(id);
ALTER TABLE wallets ADD CONSTRAINT wallets_created_by_id_foreign
	FOREIGN KEY (created_by_id) REFERENCES users(id);
ALTER TABLE wallets ADD CONSTRAINT wallets_updated_by_id_foreign
	FOREIGN KEY (updated_by_id) REFERENCES users(id);
		`)

		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping wallets")

		_, err := db.Exec(`DROP TABLE IF EXISTS wallets`)

		return err
	})
}
