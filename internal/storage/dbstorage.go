package storage

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/ole-larsen/plutonium/internal/storage/db/repository"
)

const (
	postgresProto = "postgres"
)

var NewPGSQLStorage = NewDBStorage

type DBStorageInterface interface {
	Init(ctx context.Context, sqlxDB *sqlx.DB) (*sqlx.DB, error)
	ConnectRepository(name string, sqlxDB *sqlx.DB) error
	GetUsersRepository() *repository.UsersRepository
	GetContractsRepository() *repository.ContractsRepository
	GetMenusRepository() *repository.MenusRepository
	GetSlidersRepository() *repository.SlidersRepository
	GetFilesRepository() *repository.FilesRepository
	GetCategoriesRepository() *repository.CategoriesRepository
	GetAuthorsRepository() *repository.AuthorsRepository
}

// DBStorage - database storage functionality. Use PostgreSQL version 14 or higher as a DBMS.
type DBStorage struct {
	repositories struct {
		users      repository.UsersRepositoryInterface
		contracts  repository.ContractsRepositoryInterface
		pages      repository.PagesRepositoryInterface
		categories repository.CategoriesRepositoryInterface
		menus      repository.MenusRepositoryInterface
		sliders    repository.SlidersRepositoryInterface
		files      repository.FilesRepositoryInterface
		authors    repository.AuthorsRepositoryInterface
	}
	DSN string
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

func (s *DBStorage) ConnectRepository(name string, sqlxDB *sqlx.DB) error {
	switch name {
	case "users":
		s.repositories.users = repository.NewUsersRepository(sqlxDB, name)
		if s.repositories.users.InnerDB() == nil {
			return NewError(fmt.Errorf("failed to connect %s repository", name))
		}
	case "contracts":
		s.repositories.contracts = repository.NewContractsRepository(sqlxDB, name)
		if s.repositories.users.InnerDB() == nil {
			return NewError(fmt.Errorf("failed to connect %s repository", name))
		}
	case "pages":
		s.repositories.pages = repository.NewPagesRepository(sqlxDB, name)
		if s.repositories.pages.InnerDB() == nil {
			return NewError(fmt.Errorf("failed to connect %s repository", name))
		}
	case "categories":
		s.repositories.categories = repository.NewCategoriesRepository(sqlxDB, name)
		if s.repositories.categories.InnerDB() == nil {
			return NewError(fmt.Errorf("failed to connect %s repository", name))
		}
	case "menus":
		s.repositories.menus = repository.NewMenusRepository(sqlxDB, name)
		if s.repositories.menus.InnerDB() == nil {
			return NewError(fmt.Errorf("failed to connect %s repository", name))
		}
	case "sliders":
		s.repositories.sliders = repository.NewSlidersRepository(sqlxDB, name)
		if s.repositories.sliders.InnerDB() == nil {
			return NewError(fmt.Errorf("failed to connect %s repository", name))
		}
	case "files":
		s.repositories.files = repository.NewFilesRepository(sqlxDB, name)
		if s.repositories.files.InnerDB() == nil {
			return NewError(fmt.Errorf("failed to connect %s repository", name))
		}
	case "authors":
		s.repositories.authors = repository.NewAuthorsRepository(sqlxDB, name)
		if s.repositories.files.InnerDB() == nil {
			return NewError(fmt.Errorf("failed to connect %s repository", name))
		}
	default:
		return NewError(fmt.Errorf("unknown repository: %s", name))
	}

	return nil
}

func (s *DBStorage) GetUsersRepository() *repository.UsersRepository {
	return s.repositories.users.(*repository.UsersRepository)
}

func (s *DBStorage) GetContractsRepository() *repository.ContractsRepository {
	return s.repositories.contracts.(*repository.ContractsRepository)
}

func (s *DBStorage) GetMenusRepository() *repository.MenusRepository {
	return s.repositories.menus.(*repository.MenusRepository)
}

func (s *DBStorage) GetSlidersRepository() *repository.SlidersRepository {
	return s.repositories.sliders.(*repository.SlidersRepository)
}

func (s *DBStorage) GetFilesRepository() *repository.FilesRepository {
	return s.repositories.files.(*repository.FilesRepository)
}

func (s *DBStorage) GetCategoriesRepository() *repository.CategoriesRepository {
	return s.repositories.categories.(*repository.CategoriesRepository)
}

func (s *DBStorage) GetAuthorsRepository() *repository.AuthorsRepository {
	return s.repositories.authors.(*repository.AuthorsRepository)
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

	for _, name := range []string{"users", "contracts", "pages", "categories", "menus", "sliders", "files", "authors"} {
		if err := store.ConnectRepository(name, sqlxdb); err != nil {
			return nil, err
		}
	}

	return store, nil
}
