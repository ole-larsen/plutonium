package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table menus_sections_pages")

		_, err := db.Exec(`
			CREATE TABLE IF NOT EXISTS menus_sections_pages (
				page_id INT NOT NULL,
    			section_id  INT NOT NULL,
    			order_by    INTEGER,
    			PRIMARY KEY (page_id, section_id),
    			CONSTRAINT fk_page FOREIGN KEY(page_id) REFERENCES pages(id) ON UPDATE CASCADE ON DELETE CASCADE,
    			CONSTRAINT fk_section FOREIGN KEY(section_id) REFERENCES menus_sections(id) ON UPDATE CASCADE ON DELETE CASCADE
			);
		`)

		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping menus_sections_pages")

		_, err := db.Exec(`DROP TABLE IF EXISTS menus_sections_pages`)

		return err
	})
}
