package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table pages")

		_, err := db.Exec(`
			CREATE TABLE IF NOT EXISTS pages (
				id                     SERIAL PRIMARY KEY,
				category_id            INTEGER NOT NULL,
				title                  varchar(255) UNIQUE NOT NULL,
				slug                   varchar(255),
			    description            text,
			    content                text,
			    image_id               INTEGER,
				enabled                bool NOT NULL DEFAULT TRUE,
			    is_menu                bool NOT NULL DEFAULT FALSE,
			    order_by               integer,
			    created_by_id          integer,
				updated_by_id          integer,
				created                TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
				updated                TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
				deleted                TIMESTAMP WITH TIME ZONE DEFAULT NULL
			);
			
			ALTER TABLE pages ADD CONSTRAINT pages_category_id_foreign
				FOREIGN KEY (category_id) REFERENCES categories(id) ON UPDATE CASCADE ON DELETE CASCADE;
			ALTER TABLE pages ADD CONSTRAINT pages_image_id_foreign
				FOREIGN KEY (image_id) REFERENCES files(id);
			ALTER TABLE pages ADD CONSTRAINT pages_created_by_id_foreign
				FOREIGN KEY (created_by_id) REFERENCES users(id);
			ALTER TABLE pages ADD CONSTRAINT pages_updated_by_id_foreign
				FOREIGN KEY (updated_by_id) REFERENCES users(id);
		`)

		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping pages")

		_, err := db.Exec(`DROP TABLE IF EXISTS pages`)

		return err
	})
}
