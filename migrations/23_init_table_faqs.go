package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table faqs")

		_, err := db.Exec(`
CREATE TABLE IF NOT EXISTS faqs (
	id                     SERIAL PRIMARY KEY,
	question               varchar(255) UNIQUE NOT NULL,
	answer                 text,
	order_by               integer,
	enabled                bool NOT NULL DEFAULT TRUE,
	created_by_id          integer,
	updated_by_id          integer,
	created                TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
	updated                TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
	deleted                TIMESTAMP WITH TIME ZONE DEFAULT NULL
);
ALTER TABLE faqs ADD CONSTRAINT faqs_created_by_id_foreign
	FOREIGN KEY (created_by_id) REFERENCES users(id);
ALTER TABLE faqs ADD CONSTRAINT faqs_updated_by_id_foreign
	FOREIGN KEY (updated_by_id) REFERENCES users(id);`)

		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping faqs")

		_, err := db.Exec(`DROP TABLE IF EXISTS faqs`)

		return err
	})
}
