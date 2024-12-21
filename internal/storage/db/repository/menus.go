package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-openapi/strfmt"
	"github.com/jmoiron/sqlx"
	"github.com/ole-larsen/plutonium/models"
)

type MenuSection struct {
	ID        int64       `db:"id"`
	MenuID    int64       `db:"menu_id"`
	Title     string      `db:"title"`
	Enabled   bool        `db:"enabled"`
	OrderBy   int64       `db:"order_by"`
	CreatedBy int64       `db:"created_by"`
	UpdatedBy int64       `db:"updated_by"`
	Created   strfmt.Date `db:"created"`
	Updated   strfmt.Date `db:"updated"`
	Deleted   strfmt.Date `db:"deleted"`
}

type Menu struct {
	ID        int64       `db:"id"`
	Title     string      `db:"title"`
	Enabled   bool        `db:"enabled"`
	CreatedBy int64       `db:"created_by"`
	UpdatedBy int64       `db:"updated_by"`
	Created   strfmt.Date `db:"created"`
	Updated   strfmt.Date `db:"updated"`
	Deleted   strfmt.Date `db:"deleted"`
}

type MenusRepositoryInterface interface {
	InnerDB() *sqlx.DB
	Ping() error
	MigrateContext(ctx context.Context) error

	GetMenuByProvider(provider string) (*models.PublicMenu, error)
}

// MenusRepository - repository to store frontend menus.
type MenusRepository struct {
	DB  sqlx.DB
	TBL string
}

func NewMenusRepository(db *sqlx.DB, tbl string) *MenusRepository {
	if db == nil {
		return nil
	}

	return &MenusRepository{
		DB:  *db,
		TBL: tbl,
	}
}

func (r *MenusRepository) InnerDB() *sqlx.DB {
	if r == nil {
		return nil
	}

	return &r.DB
}

func (r *MenusRepository) Ping() error {
	if r == nil {
		return ErrDBNotInitialized
	}

	return r.DB.Ping()
}

func (r *MenusRepository) MigrateContext(ctx context.Context) error {
	if r == nil {
		return ErrDBNotInitialized
	}

	_, err := r.DB.ExecContext(ctx, fmt.Sprintf(`
			CREATE TABLE IF NOT EXISTS %s (
				id                     SERIAL PRIMARY KEY,
				title                  varchar(255),
			    enabled                bool NOT NULL DEFAULT TRUE,
			    created_by_id          integer,
				updated_by_id          integer,
			    created                TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
				updated                TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
				deleted                TIMESTAMP WITH TIME ZONE DEFAULT NULL
			);
			
			CREATE TABLE IF NOT EXISTS menus_sections (
				id                     SERIAL PRIMARY KEY,
				menu_id                integer,
				title                  varchar(255),
			    enabled                bool NOT NULL DEFAULT TRUE,
			    order_by               integer,
			    created_by_id          integer,
				updated_by_id          integer,
			    created                TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
				updated                TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
				deleted                TIMESTAMP WITH TIME ZONE DEFAULT NULL
			);
			
			ALTER TABLE menus_sections ADD CONSTRAINT menus_sections_menu_id_foreign
				FOREIGN KEY (menu_id) REFERENCES menus(id) ON UPDATE CASCADE ON DELETE CASCADE;

			CREATE TABLE IF NOT EXISTS menus_sections_pages (
				page_id INT NOT NULL,
    			section_id  INT NOT NULL,
    			order_by    INTEGER,
    			PRIMARY KEY (page_id, section_id),
    			CONSTRAINT fk_page FOREIGN KEY(page_id) REFERENCES pages(id) ON UPDATE CASCADE ON DELETE CASCADE,
    			CONSTRAINT fk_section FOREIGN KEY(section_id) REFERENCES menus_sections(id) ON UPDATE CASCADE ON DELETE CASCADE
			);
			`, r.TBL))

	return err
}

func (r *MenusRepository) GetMenuByProvider(provider string) (*models.PublicMenu, error) {
	if r == nil {
		return nil, ErrDBNotInitialized
	}
	var menu models.PublicMenu
	var attributes AggregatedMenuAttributesJSON

	sqlStatement := `SELECT
		m1.id,
		JSON_BUILD_OBJECT(
			'name',     m1.title,
			'items',    (SELECT ARRAY_AGG(JSON_BUILD_OBJECT(
				'id', ms.id,
				'attributes', (SELECT JSON_BUILD_OBJECT(
					'name',     ms.title,
					'orderBy', ms.order_by,
					'items',    (SELECT ARRAY_AGG(JSON_BUILD_OBJECT(
									'id', p.id,
									'attributes', (SELECT JSON_BUILD_OBJECT(
										'name',     p.title,
										'link',     p.slug,
										'orderBy', msp.order_by
									) FROM pages p2 WHERE p2.id = p.id)
								) ORDER BY msp.order_by ASC) 
								FROM menus_sections_pages msp JOIN pages p ON (msp.page_id = p.id AND p.enabled = TRUE AND p.deleted_at isNull) WHERE msp.section_id = ms2.id)
				) FROM menus_sections ms2 WHERE ms.id = ms2.id AND ms.enabled = TRUE)
			) ORDER BY ms.order_by ASC) FROM menus_sections ms WHERE ms.menu_id = m1.id AND ms.enabled = TRUE)
		) as attributes
	FROM menus m1
	WHERE m1.title = $1 AND m1.enabled = TRUE AND m1.deleted_at isNull;`
	row := r.DB.QueryRow(sqlStatement, provider)
	err := row.Scan(&menu.ID, &attributes)
	switch err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		menu.Attributes = &models.PublicMenuAttributes{
			Items: attributes.Items,
			Link:  attributes.Link,
			Name:  attributes.Name,
		}
		return &menu, nil
	default:
		return nil, err
	}
}
