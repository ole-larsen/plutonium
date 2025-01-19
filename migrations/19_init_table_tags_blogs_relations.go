package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating tags blogs relations")

		_, err := db.Exec(`
			CREATE TABLE tags_blogs (
  		  		tag_id INT NOT NULL,
    			blog_id INT NOT NULL,
    			PRIMARY KEY (tag_id, blog_id),
    			CONSTRAINT fk_tag FOREIGN KEY(tag_id) REFERENCES tags(id) ON UPDATE CASCADE ON DELETE CASCADE,
    			CONSTRAINT fk_blog FOREIGN KEY(blog_id) REFERENCES blogs(id) ON UPDATE CASCADE ON DELETE CASCADE
			);
		`)

		return err
	}, func(db migrations.DB) error {
		fmt.Println("removing tag blogs relations")

		_, err := db.Exec(`DROP TABLE IF EXISTS tags_blogs;`)

		return err
	})
}
