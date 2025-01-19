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
	GetPagesRepository() *repository.PagesRepository
	GetContactsRepository() *repository.ContactsRepository
	GetFaqsRepository() *repository.FaqsRepository
	GetContactFormsRepository() *repository.ContactFormRepository
	GetHelpCenterRepository() *repository.HelpCenterRepository
	GetTagsRepository() *repository.TagsRepository
	GetBlogsRepository() *repository.BlogsRepository
}

// DBStorage - database storage functionality. Use PostgreSQL version 14 or higher as a DBMS.
type DBStorage struct {
	repositories struct {
		users        repository.UsersRepositoryInterface
		contracts    repository.ContractsRepositoryInterface
		pages        repository.PagesRepositoryInterface
		categories   repository.CategoriesRepositoryInterface
		menus        repository.MenusRepositoryInterface
		sliders      repository.SlidersRepositoryInterface
		files        repository.FilesRepositoryInterface
		authors      repository.AuthorsRepositoryInterface
		contacts     repository.ContactsRepositoryInterface
		contactForms repository.ContactFormRepositoryInterface
		faqs         repository.FaqsRepositoryInterface
		helpCenter   repository.HelpCenterRepositoryInterface
		tags         repository.TagsRepositoryInterface
		blogs        repository.BlogsRepositoryInterface
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
	case "contacts":
		s.repositories.contacts = repository.NewContactsRepository(sqlxDB, name)
		if s.repositories.contacts.InnerDB() == nil {
			return NewError(fmt.Errorf("failed to connect %s repository", name))
		}
	case "contactForms":
		s.repositories.contactForms = repository.NewContactFormRepository(sqlxDB, name)
		if s.repositories.contactForms.InnerDB() == nil {
			return NewError(fmt.Errorf("failed to connect %s repository", name))
		}
	case "faqs":
		s.repositories.faqs = repository.NewFaqsRepository(sqlxDB, name)
		if s.repositories.faqs.InnerDB() == nil {
			return NewError(fmt.Errorf("failed to connect %s repository", name))
		}
	case "helpCenter":
		s.repositories.helpCenter = repository.NewHelpCenterRepository(sqlxDB, name)
		if s.repositories.helpCenter.InnerDB() == nil {
			return NewError(fmt.Errorf("failed to connect %s repository", name))
		}
	case "blogs":
		s.repositories.blogs = repository.NewBlogsRepository(sqlxDB, name)
		if s.repositories.blogs.InnerDB() == nil {
			return NewError(fmt.Errorf("failed to connect %s repository", name))
		}
	case "tags":
		s.repositories.tags = repository.NewTagsRepository(sqlxDB, name, s.GetBlogsRepository(), s.GetPagesRepository())
		if s.repositories.tags.InnerDB() == nil {
			return NewError(fmt.Errorf("failed to connect %s repository", name))
		}
	default:
		return NewError(fmt.Errorf("unknown repository: %s", name))
	}

	return nil
}

func (s *DBStorage) GetUsersRepository() *repository.UsersRepository {
	if repo, ok := s.repositories.users.(*repository.UsersRepository); ok {
		return repo
	}

	return nil
}

func (s *DBStorage) GetContractsRepository() *repository.ContractsRepository {
	if repo, ok := s.repositories.contracts.(*repository.ContractsRepository); ok {
		return repo
	}

	return nil
}

func (s *DBStorage) GetMenusRepository() *repository.MenusRepository {
	if repo, ok := s.repositories.menus.(*repository.MenusRepository); ok {
		return repo
	}

	return nil
}

func (s *DBStorage) GetSlidersRepository() *repository.SlidersRepository {
	if repo, ok := s.repositories.sliders.(*repository.SlidersRepository); ok {
		return repo
	}

	return nil
}

func (s *DBStorage) GetFilesRepository() *repository.FilesRepository {
	if repo, ok := s.repositories.files.(*repository.FilesRepository); ok {
		return repo
	}

	return nil
}

func (s *DBStorage) GetCategoriesRepository() *repository.CategoriesRepository {
	if repo, ok := s.repositories.categories.(*repository.CategoriesRepository); ok {
		return repo
	}

	return nil
}

func (s *DBStorage) GetAuthorsRepository() *repository.AuthorsRepository {
	if repo, ok := s.repositories.authors.(*repository.AuthorsRepository); ok {
		return repo
	}

	return nil
}

func (s *DBStorage) GetPagesRepository() *repository.PagesRepository {
	if repo, ok := s.repositories.pages.(*repository.PagesRepository); ok {
		return repo
	}

	return nil
}

func (s *DBStorage) GetContactsRepository() *repository.ContactsRepository {
	if repo, ok := s.repositories.contacts.(*repository.ContactsRepository); ok {
		return repo
	}

	return nil
}

func (s *DBStorage) GetContactFormsRepository() *repository.ContactFormRepository {
	if repo, ok := s.repositories.contactForms.(*repository.ContactFormRepository); ok {
		return repo
	}

	return nil
}

func (s *DBStorage) GetFaqsRepository() *repository.FaqsRepository {
	if repo, ok := s.repositories.faqs.(*repository.FaqsRepository); ok {
		return repo
	}

	return nil
}

func (s *DBStorage) GetHelpCenterRepository() *repository.HelpCenterRepository {
	if repo, ok := s.repositories.helpCenter.(*repository.HelpCenterRepository); ok {
		return repo
	}

	return nil
}

func (s *DBStorage) GetTagsRepository() *repository.TagsRepository {
	if repo, ok := s.repositories.tags.(*repository.TagsRepository); ok {
		return repo
	}

	return nil
}

func (s *DBStorage) GetBlogsRepository() *repository.BlogsRepository {
	if repo, ok := s.repositories.blogs.(*repository.BlogsRepository); ok {
		return repo
	}

	return nil
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

	for _, name := range []string{
		"users",
		"contracts",
		"pages",
		"categories",
		"menus",
		"sliders",
		"files",
		"authors",
		"contacts",
		"contactForms",
		"faqs",
		"helpCenter",
		"blogs",
		"tags",
	} {
		if err := store.ConnectRepository(name, sqlxdb); err != nil {
			return nil, err
		}
	}

	return store, nil
}
