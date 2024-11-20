// Package storage contains all types of storages using by client and server. any storage type can be replaced by any other from this package.
// Copyright 2024 The Oleg Nazarov. All rights reserved.
package storage

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/ole-larsen/plutonium/internal/storage/db/repository"
)

// Storage - main interface for all types of storage.
type Storage interface {
	Init(ctx context.Context, sqlxDB *sqlx.DB) (*sqlx.DB, error)
	Ping() error
	ConnectUsersRepository(ctx context.Context, sqlxDB *sqlx.DB) error
	CreateUser(ctx context.Context, userMap map[string]interface{}) error
	GetUser(ctx context.Context, email string) (*repository.User, error)
}
