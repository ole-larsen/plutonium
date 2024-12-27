package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("insert into table roles")

		_, err := db.Exec(`
INSERT INTO roles (title, description, enabled) VALUES ( 'superadmin', 'main system role', true) ON CONFLICT (title) DO NOTHING;
INSERT INTO roles (title, description, enabled) VALUES ( 'admin',      'admin system role', true) ON CONFLICT (title) DO NOTHING;
INSERT INTO roles (title, description, enabled) VALUES ( 'user',       'default user role', true) ON CONFLICT (title) DO NOTHING;
INSERT INTO roles (title, description, enabled) VALUES ( 'manager',    'default manager role', true) ON CONFLICT (title) DO NOTHING;
		`)

		return err
	}, func(db migrations.DB) error {
		fmt.Println("remove from table roles")

		_, err := db.Exec(`DELETE FROM roles`)

		return err
	})
}
