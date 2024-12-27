package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table files")

		_, err := db.Exec(`
CREATE TABLE IF NOT EXISTS files (
	id SERIAL PRIMARY KEY,
	name                VARCHAR(255) UNIQUE NOT NULL,
	alt                 VARCHAR(128),
	caption             VARCHAR(255),
	width               integer,
	height              integer,
	formats             jsonb,
	hash                VARCHAR(128),
	ext                 VARCHAR(6),
	mime                VARCHAR(12),
	size                integer,
	url                 VARCHAR(255),
	blur                VARCHAR(255),
	preview_url         VARCHAR(255),
	provider            VARCHAR(128),
	provider_metadata   jsonb,
	created_by_id       integer,
	updated_by_id       integer,
	created             TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
	updated             TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
	deleted             TIMESTAMP WITH TIME ZONE DEFAULT NULL
);
ALTER TABLE files ADD CONSTRAINT files_created_by_id_foreign
	FOREIGN KEY (created_by_id) REFERENCES users(id);
ALTER TABLE files ADD CONSTRAINT files_updated_by_id_foreign
	FOREIGN KEY (updated_by_id) REFERENCES users(id);`)

		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping files")

		_, err := db.Exec(`DROP TABLE IF EXISTS files`)

		return err
	})
}
