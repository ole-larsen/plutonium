package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table sliders items")

		_, err := db.Exec(`
CREATE TABLE IF NOT EXISTS sliders_items (
	id                     SERIAL PRIMARY KEY,
	slider_id              INTEGER NOT NULL,
	heading                text,
	description            text,
	btn_link_1             varchar(128),
	btn_text_1             varchar(128),
	btn_link_2             varchar(128),
	btn_text_2             varchar(128),
	image_id               INTEGER,
	bg_image_id            INTEGER,
	enabled                bool NOT NULL DEFAULT TRUE,
	created_by_id          integer,
	updated_by_id          integer,
	created                TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
	updated                TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
	deleted                TIMESTAMP WITH TIME ZONE DEFAULT NULL
);
ALTER TABLE sliders_items ADD CONSTRAINT sliders_items_slider_id_foreign
	FOREIGN KEY (slider_id) REFERENCES sliders(id) ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE sliders_items ADD CONSTRAINT sliders_items_image_id_foreign
	FOREIGN KEY (image_id) REFERENCES files(id);
ALTER TABLE sliders_items ADD CONSTRAINT sliders_items_bg_image_id_foreign
	FOREIGN KEY (bg_image_id) REFERENCES files(id);
ALTER TABLE sliders_items ADD CONSTRAINT sliders_items_created_by_id_foreign
	FOREIGN KEY (created_by_id) REFERENCES users(id);
ALTER TABLE sliders_items ADD CONSTRAINT sliders_items_updated_by_id_foreign
	FOREIGN KEY (updated_by_id) REFERENCES users(id);
		`)

		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping sliders items")

		_, err := db.Exec(`DROP TABLE IF EXISTS sliders_items`)

		return err
	})
}
