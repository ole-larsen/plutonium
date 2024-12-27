package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table menus_sections")

		_, err := db.Exec(`
CREATE TABLE IF NOT EXISTS menus_sections (
	id                     SERIAL PRIMARY KEY,
	menu_id                integer,
	title                  varchar(255),
	enabled                bool NOT NULL DEFAULT TRUE,
	order_by               integer,
	created_by_id          integer,
	updated_by_id          integer,
	created                TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
	updated                TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
	deleted                TIMESTAMP WITH TIME ZONE DEFAULT NULL
);

ALTER TABLE menus_sections ADD CONSTRAINT menus_sections_menu_id_foreign
	FOREIGN KEY (menu_id) REFERENCES menus(id) ON UPDATE CASCADE ON DELETE CASCADE;
		`)

		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping menus_sections")

		_, err := db.Exec(`DROP TABLE IF EXISTS menus_sections`)

		return err
	})
}
