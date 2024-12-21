package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table sliders")

		_, err := db.Exec(`
			CREATE TABLE IF NOT EXISTS sliders (
				id                     SERIAL PRIMARY KEY,
				provider               varchar(255),
				title                  varchar(255) UNIQUE NOT NULL,
				description            text NOT NULL,
				enabled                bool NOT NULL DEFAULT TRUE,
			    created_by_id          integer,
				updated_by_id          integer,
				created                TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
				updated                TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
				deleted                TIMESTAMP WITH TIME ZONE DEFAULT NULL
			);
			ALTER TABLE sliders ADD CONSTRAINT sliders_created_by_id_foreign
				FOREIGN KEY (created_by_id) REFERENCES users(id);
			ALTER TABLE sliders ADD CONSTRAINT sliders_updated_by_id_foreign
				FOREIGN KEY (updated_by_id) REFERENCES users(id);
	
			`)

		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping sliders")

		_, err := db.Exec(`DROP TABLE IF EXISTS sliders`)

		return err
	})
}
