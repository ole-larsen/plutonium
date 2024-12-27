package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating users_roles relations")

		_, err := db.Exec(`
CREATE TABLE users_roles (
	user_id INT NOT NULL,
	role_id INT NOT NULL DEFAULT 3,
	PRIMARY KEY (user_id, role_id),
	CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE,
	CONSTRAINT fk_role FOREIGN KEY(role_id) REFERENCES roles(id) ON UPDATE CASCADE ON DELETE CASCADE
);
		`)

		return err
	}, func(db migrations.DB) error {
		fmt.Println("removing users_roles relations")

		_, err := db.Exec(`DROP TABLE IF EXISTS users_roles;`)

		return err
	})
}
