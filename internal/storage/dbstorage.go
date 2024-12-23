package storage

import (
	"context"
	"fmt"

	"github.com/dgryski/dgoogauth"
	"github.com/jmoiron/sqlx"

	"github.com/ole-larsen/plutonium/internal/otp"
	"github.com/ole-larsen/plutonium/internal/storage/db/repository"
)

const (
	postgresProto = "postgres"
)

var NewPGSQLStorage = NewDBStorage

type DBStorageInterface interface {
	Init(ctx context.Context, sqlxDB *sqlx.DB) (*sqlx.DB, error)
	Ping() error
	ConnectUsersRepository(ctx context.Context, sqlxDB *sqlx.DB) error
	ConnectContractsRepository(ctx context.Context, sqlxDB *sqlx.DB) error
	ConnectPagesRepository(ctx context.Context, sqlxDB *sqlx.DB) error
	ConnectCategoriesRepository(ctx context.Context, sqlxDB *sqlx.DB) error
	ConnectMenusRepository(ctx context.Context, sqlxDB *sqlx.DB) error
	CreateUser(ctx context.Context, userMap map[string]interface{}) (*dgoogauth.OTPConfig, error)
	GetUser(ctx context.Context, login string) (*repository.User, error)
	GetContractsRepository() *repository.ContractsRepository
	GetMenusRepository() *repository.MenusRepository
}

// DBStorage - database storage functionality. Use PostgreSQL version 14 or higher as a DBMS.
type DBStorage struct {
	Users      repository.UsersRepositoryInterface
	Contracts  *repository.ContractsRepository
	Pages      *repository.PagesRepository
	Categories *repository.CategoriesRepository
	Menus      *repository.MenusRepository
	DSN        string
}

func NewDBStorage(dsn string) DBStorageInterface {
	if dsn == "" {
		return nil
	}

	return &DBStorage{
		DSN: dsn,
	}
}

// Init connects to the database using the DSN provided in DBStorage.
func (s *DBStorage) Init(ctx context.Context, sqlxDB *sqlx.DB) (*sqlx.DB, error) {
	if s == nil {
		return nil, NewError(fmt.Errorf("DBStorage is nil"))
	}
	// for test cases pass sqlxDb as nil
	if sqlxDB == nil {
		var err error

		sqlxDB, err = sqlx.ConnectContext(ctx, postgresProto, s.DSN)
		if err != nil {
			return nil, NewError(fmt.Errorf("failed to connect to the database: %w", err))
		}
	}

	return sqlxDB, nil
}

func (s *DBStorage) ConnectUsersRepository(_ context.Context, sqlxDB *sqlx.DB) error {
	s.Users = repository.NewUsersRepository(sqlxDB, "users")
	if s.Users.InnerDB() == nil {
		return NewError(fmt.Errorf("failed to connect users repository"))
	}

	return nil
}

func (s *DBStorage) ConnectContractsRepository(_ context.Context, sqlxDB *sqlx.DB) error {
	s.Contracts = repository.NewContractsRepository(sqlxDB, "contracts")
	if s.Contracts.InnerDB() == nil {
		return NewError(fmt.Errorf("failed to connect contracts repository"))
	}

	return nil
}

func (s *DBStorage) ConnectPagesRepository(_ context.Context, sqlxDB *sqlx.DB) error {
	s.Pages = repository.NewPagesRepository(sqlxDB, "pages")
	if s.Pages.InnerDB() == nil {
		return NewError(fmt.Errorf("failed to connect pages repository"))
	}

	return nil
}

func (s *DBStorage) ConnectMenusRepository(_ context.Context, sqlxDB *sqlx.DB) error {
	s.Menus = repository.NewMenusRepository(sqlxDB, "menus")
	if s.Menus.InnerDB() == nil {
		return NewError(fmt.Errorf("failed to connect menus repository"))
	}

	return nil
}

func (s *DBStorage) ConnectCategoriesRepository(_ context.Context, sqlxDB *sqlx.DB) error {
	s.Categories = repository.NewCategoriesRepository(sqlxDB, "categories")
	if s.Categories.InnerDB() == nil {
		return NewError(fmt.Errorf("failed to connect categories repository"))
	}

	return nil
}

func (s *DBStorage) Ping() error {
	if s == nil || s.Users == nil || s.Users.InnerDB() == nil {
		return NewError(fmt.Errorf("DBStorage is nil or not initialized"))
	}

	return s.Users.Ping()
}

func (s *DBStorage) CreateUser(ctx context.Context, userMap map[string]interface{}) (*dgoogauth.OTPConfig, error) {
	if s == nil || s.Users == nil || s.Users.InnerDB() == nil {
		return nil, NewError(fmt.Errorf("DBStorage is nil or not initialized"))
	}

	otpc := otp.CreateOTP()
	userMap["secret"] = otpc.Secret

	return otpc, s.Users.Create(ctx, userMap)
}

func (s *DBStorage) GetUser(ctx context.Context, login string) (*repository.User, error) {
	if s == nil || s.Users == nil || s.Users.InnerDB() == nil {
		return nil, NewError(fmt.Errorf("DBStorage is nil or not initialized"))
	}

	return s.Users.GetOne(ctx, login)
}

func (s *DBStorage) GetContractsRepository() *repository.ContractsRepository {
	return s.Contracts
}

func (s *DBStorage) GetMenusRepository() *repository.MenusRepository {
	return s.Menus
}

func SetupStorage(ctx context.Context, dsn string) (DBStorageInterface, error) {
	store := NewPGSQLStorage(dsn)
	if store == nil {
		return nil, NewError(fmt.Errorf("cannot init db using dsn='%s'", dsn))
	}

	sqlxdb, err := store.Init(ctx, nil)
	if err != nil {
		return nil, err
	}

	if err := store.ConnectUsersRepository(ctx, sqlxdb); err != nil {
		return nil, err
	}

	if err := store.ConnectContractsRepository(ctx, sqlxdb); err != nil {
		return nil, err
	}

	if err := store.ConnectPagesRepository(ctx, sqlxdb); err != nil {
		return nil, err
	}

	if err := store.ConnectCategoriesRepository(ctx, sqlxdb); err != nil {
		return nil, err
	}

	if err := store.ConnectMenusRepository(ctx, sqlxdb); err != nil {
		return nil, err
	}

	return store, nil
}
