package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table users")

		_, err := db.Exec(`
			CREATE TABLE IF NOT EXISTS users (
			    id                     SERIAL PRIMARY KEY,
				uuid                   varchar(255) NOT NULL DEFAULT gen_random_uuid(),
			    username               varchar(255) UNIQUE NOT NULL,
				address                varchar(64)[],
				email                  varchar(255) UNIQUE NOT NULL,
			    password               varchar(255) NOT NULL,
				password_reset_token   varchar(255),
				password_reset_expires BIGINT,
				nonce                  varchar(255),
				rsa_secret             varchar(255),
				secret                 varchar(255),
				gravatar               varchar(255),
				wallpaper_id           INTEGER,
			    enabled                bool NOT NULL DEFAULT FALSE,
				created                TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
				updated                TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
				deleted                TIMESTAMP WITH TIME ZONE DEFAULT NULL,
				CHECK (LENGTH(username) > 4),
			    CHECK (LENGTH(password) > 4)
			);
			CREATE INDEX idx_users_address ON users USING GIN(address);
		`)

		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping table users")

		_, err := db.Exec(`DROP TABLE IF EXISTS users; DROP INDEX if EXISTS idx_users_address`)

		return err
	})
}
