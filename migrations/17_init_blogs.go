package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table blogs")
		_, err := db.Exec(`
			CREATE TABLE IF NOT EXISTS blogs (
				id                     SERIAL PRIMARY KEY,
				title                  varchar(255) UNIQUE NOT NULL,
				slug                   varchar(255),
			    quote                  text,
			    description            jsonb,
			    content                text,
			    author_id              INTEGER NOT NULL,
			    image_id               INTEGER,
			    image_1_id             INTEGER,
			    image_2_id             INTEGER,
			    image_3_id             INTEGER,
			    public_date            TIMESTAMP WITH TIME ZONE,
				enabled                bool NOT NULL DEFAULT TRUE,
			    order_by               integer,
			    created_by_id          integer,
				updated_by_id          integer,
				created                TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
				updated                TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
				deleted                TIMESTAMP WITH TIME ZONE DEFAULT NULL
			);
			
			ALTER TABLE blogs ADD CONSTRAINT blogs_author_id_foreign
				FOREIGN KEY (author_id) REFERENCES authors(id) ON UPDATE CASCADE ON DELETE CASCADE;
			
			ALTER TABLE blogs ADD CONSTRAINT blogs_image_id_foreign
				FOREIGN KEY (image_id) REFERENCES files(id);
		
			ALTER TABLE blogs ADD CONSTRAINT blogs_image_1_id_foreign
				FOREIGN KEY (image_1_id) REFERENCES files(id);

			ALTER TABLE blogs ADD CONSTRAINT blogs_image_2_id_foreign
				FOREIGN KEY (image_2_id) REFERENCES files(id);

			ALTER TABLE blogs ADD CONSTRAINT blogs_image_3_id_foreign
				FOREIGN KEY (image_3_id) REFERENCES files(id);
		`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping pages")
		_, err := db.Exec(`DROP TABLE IF EXISTS blogs`)
		return err
	})
}
