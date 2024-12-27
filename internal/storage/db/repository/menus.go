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
	Created   strfmt.Date `db:"created"`
	Updated   strfmt.Date `db:"updated"`
	Deleted   strfmt.Date `db:"deleted"`
	Title     string      `db:"title"`
	ID        int64       `db:"id"`
	MenuID    int64       `db:"menu_id"`
	OrderBy   int64       `db:"order_by"`
	CreatedBy int64       `db:"created_by"`
	UpdatedBy int64       `db:"updated_by"`
	Enabled   bool        `db:"enabled"`
}

type Menu struct {
	Created   strfmt.Date `db:"created"`
	Updated   strfmt.Date `db:"updated"`
	Deleted   strfmt.Date `db:"deleted"`
	Title     string      `db:"title"`
	ID        int64       `db:"id"`
	CreatedBy int64       `db:"created_by"`
	UpdatedBy int64       `db:"updated_by"`
	Enabled   bool        `db:"enabled"`
}

type MenusRepositoryInterface interface {
	InnerDB() *sqlx.DB
	Ping() error

	GetMenuByProvider(ctx context.Context, provider string) (*models.PublicMenu, error)
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

func (r *MenusRepository) GetMenuByProvider(ctx context.Context, provider string) (*models.PublicMenu, error) {
	if r == nil {
		return nil, ErrDBNotInitialized
	}

	var menu models.PublicMenu

	var attributes AggregatedMenuAttributesJSON

	sqlStatement := `
SELECT
	m1.id,
	JSON_BUILD_OBJECT(
		'name',     m1.title,
		'items',    (SELECT ARRAY_AGG(JSON_BUILD_OBJECT(
			'id', ms.id,
			'attributes', (SELECT JSON_BUILD_OBJECT(
				'name',     ms.title,
				'orderBy',  ms.order_by,
				'items',    (SELECT ARRAY_AGG(JSON_BUILD_OBJECT(
								'id', p.id,
								'attributes', (SELECT JSON_BUILD_OBJECT(
									'name',     p.title,
									'link',     p.slug,
									'orderBy', msp.order_by
								) FROM pages p2 WHERE p2.id = p.id)
							) ORDER BY msp.order_by ASC) 
							FROM menus_sections_pages msp JOIN pages p ON (msp.page_id = p.id AND p.enabled = TRUE AND p.deleted isNull) WHERE msp.section_id = ms2.id)
			) FROM menus_sections ms2 WHERE ms.id = ms2.id AND ms.enabled = TRUE)
		) ORDER BY ms.order_by ASC) FROM menus_sections ms WHERE ms.menu_id = m1.id AND ms.enabled = TRUE)
	) as attributes
FROM menus m1
WHERE m1.title = $1 AND m1.enabled = TRUE AND m1.deleted isNull;`

	row := r.DB.QueryRowContext(ctx, sqlStatement, provider)

	err := row.Scan(&menu.ID, &attributes)
	switch err {
	case sql.ErrNoRows:
		return nil, NewError(fmt.Errorf("[%s] menu not found", provider))
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
