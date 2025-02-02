package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table create_and_sell")

		_, err := db.Exec(`
CREATE TABLE IF NOT EXISTS create_and_sell (
	id                     SERIAL PRIMARY KEY,
	title                  varchar(128) UNIQUE NOT NULL,
	link                   varchar(128),
	image_id               integer,
	description            text,
	enabled                bool NOT NULL DEFAULT TRUE,
	order_by               integer,
	created_by_id          integer,
	updated_by_id          integer,
	created                TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
	updated                TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
	deleted                TIMESTAMP WITH TIME ZONE DEFAULT NULL
);
ALTER TABLE create_and_sell ADD CONSTRAINT create_and_sell_image_id_foreign
	FOREIGN KEY (image_id) REFERENCES files(id);
ALTER TABLE create_and_sell ADD CONSTRAINT create_and_sell_created_by_id_foreign
	FOREIGN KEY (created_by_id) REFERENCES users(id);
ALTER TABLE create_and_sell ADD CONSTRAINT create_and_sell_updated_by_id_foreign
	FOREIGN KEY (updated_by_id) REFERENCES users(id);`)

		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping create_and_sell")

		_, err := db.Exec(`DROP TABLE IF EXISTS create_and_sell`)

		return err
	})
}
