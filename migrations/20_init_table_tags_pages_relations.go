package main

import (
	"fmt"
	"github.com/go-pg/migrations"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating tags pages relations")
		_, err := db.Exec(`
			CREATE TABLE tags_pages (
  		  		tag_id INT NOT NULL,
    			page_id INT NOT NULL,
    			PRIMARY KEY (tag_id, page_id),
    			CONSTRAINT fk_tag FOREIGN KEY(tag_id) REFERENCES tags(id) ON UPDATE CASCADE ON DELETE CASCADE,
    			CONSTRAINT fk_page FOREIGN KEY(page_id) REFERENCES pages(id) ON UPDATE CASCADE ON DELETE CASCADE
			);
		`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("removing tag pages relations")
		_, err := db.Exec(`DROP TABLE IF EXISTS tags_pages;`)
		return err
	})
}
