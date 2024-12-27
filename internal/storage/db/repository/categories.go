package repository

import (
	"context"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hashicorp/go-multierror"
	"github.com/jmoiron/sqlx"
	"github.com/ole-larsen/plutonium/models"
)

type CategoryCollectionCollectible struct {
	ID         int64 `db:"id"`
	Attributes CategoryCollectionCollectibleAttributes
}

type CategoryCollectionCollectibleAttributes struct {
	Item       int64               `db:"item"`
	Tokens     []int64             `db:"tokens"`
	Collection int64               `db:"collection"`
	URI        string              `db:"uri"`
	Creator    int64               `db:"creator"`
	Owner      int64               `db:"owner"`
	Created    time.Time           `db:"created"`
	Metadata   CollectibleMetadata `db:"metadata"`
	Details    CollectibleDetails  `db:"details"`
}

type CategoryCollectionAttributes struct {
	CategoryID   int64                                         `db:"category_id"`
	Name         string                                        `db:"name"`
	Symbol       string                                        `db:"symbol"`
	Description  string                                        `db:"description"`
	Fee          string                                        `db:"fee"`
	Creator      int64                                         `db:"creator"`
	MaxItems     int64                                         `db:"max_items"`
	Owner        int64                                         `db:"owner"`
	Address      string                                        `db:"address"`
	Slug         string                                        `db:"slug"`
	URL          string                                        `db:"url"`
	Logo         AggregatedImageJSON                           `db:"logo"`
	Featured     AggregatedImageJSON                           `db:"featured"`
	Banner       AggregatedImageJSON                           `db:"banner"`
	Collectibles AggregatedCategoryCollectionsCollectiblesJSON `db:"collectibles"`
	isApproved   bool                                          `db:"is_approved"`
	isLocked     bool                                          `db:"is_locked"`
	Created      string                                        `db:"created"`
	Updated      time.Time                                     `db:"updated_at"`
	Deleted      time.Time                                     `db:"deleted_at"`
}

type CategoryCollection struct {
	ID         int64 `db:"id"`
	Attributes CategoryCollectionAttributes
}

type Category struct {
	ID          int64                             `db:"id"`
	ParentID    *int64                            `db:"parent_id"`
	Provider    string                            `db:"provider"`
	Title       string                            `db:"title"`
	Slug        string                            `db:"slug"`
	Description *string                           `db:"description"`
	Content     *string                           `db:"content"`
	ImageID     *int64                            `db:"image_id"`
	Enabled     bool                              `db:"enabled"`
	OrderBy     *int64                            `db:"order_by"`
	CreatedByID *int64                            `db:"created_by_id"`
	UpdatedByID *int64                            `db:"updated_by_id"`
	Created     time.Time                         `db:"created"`
	Updated     time.Time                         `db:"updated"`
	Deleted     time.Time                         `db:"deleted"`
	Collections AggregatedCategoryCollectionsJSON `db:"collections"`
}

type CategoriesRepositoryInterface interface {
	InnerDB() *sqlx.DB
	Ping() error
	GetPublicCollectibleCategories(ctx context.Context, users *UsersRepository) ([]*models.PublicCategory, error)
}

// CategoriesRepository - repository to store frontend categories.
type CategoriesRepository struct {
	DB  sqlx.DB
	TBL string
}

func NewCategoriesRepository(db *sqlx.DB, tbl string) *CategoriesRepository {
	if db == nil {
		return nil
	}

	return &CategoriesRepository{
		DB:  *db,
		TBL: tbl,
	}
}

func (r *CategoriesRepository) InnerDB() *sqlx.DB {
	if r == nil {
		return nil
	}

	return &r.DB
}

func (r *CategoriesRepository) Ping() error {
	if r == nil {
		return ErrDBNotInitialized
	}

	return r.DB.Ping()
}

func (r *CategoriesRepository) GetPublicCollectibleCategories(ctx context.Context, users *UsersRepository) ([]*models.PublicCategory, error) {
	if r == nil {
		return nil, ErrDBNotInitialized
	}
	var (
		multierr   multierror.Error
		categories []*models.PublicCategory
	)

	rows, err := r.DB.QueryxContext(ctx, `
SELECT
	c.id,
	c.title,
	c.slug,
	c.description,
	c.content,
	c.created,
	(SELECT JSON_BUILD_OBJECT(
		'id', f.id,
		'attributes', (SELECT JSON_BUILD_OBJECT(
			'name',            f.name,
			'alt',             f.alt,
			'caption',         f.caption,
			'ext',             f.ext,
			'provider',        f.provider,
			'width',           f.width,
			'height',          f.height,
			'size',            f.size,
			'url',             f.url
		) FROM files f WHERE f.id = c.image_id)
	) FROM files f WHERE f.id = c.image_id) as image,
	(SELECT (SELECT JSON_AGG(JSON_BUILD_OBJECT(
	'id', col.id,
	'attributes', (SELECT JSON_BUILD_OBJECT(
		'name',        col.name,
		'symbol',      col.symbol,
		'description', col.description,
		'address',     col.address,
		'fee',         col.fee,
		'max_items',   col.max_items,
		'is_approved', col.is_approved,
		'is_locked',   col.is_locked,
		'url',         col.url,
		'slug',        col.slug,
		'owner',       col.owner_id,
		'creator',     col.creator_id,
		'created',     to_char(col.created, 'Month dd, yyyy'),
		'updated',     col.updated,
		'logo',        (SELECT JSON_BUILD_OBJECT(
		'id', f.id,
		'attributes', (SELECT JSON_BUILD_OBJECT(
			'name',            f.name,
			'alt',             f.alt,
			'caption',         f.caption,
			'ext',             f.ext,
			'provider',        f.provider,
			'width',           f.width,
			'height',          f.height,
			'size',            f.size,
			'url',             f.url
		) FROM files f WHERE f.id = c1.logo_id)
		) FROM files f WHERE f.id = c1.logo_id),
		'featured',    (SELECT JSON_BUILD_OBJECT(
		'id', f.id,
		'attributes', (SELECT JSON_BUILD_OBJECT(
			'name',            f.name,
			'alt',             f.alt,
			'caption',         f.caption,
			'ext',             f.ext,
			'provider',        f.provider,
			'width',           f.width,
			'height',          f.height,
			'size',            f.size,
			'url',             f.url
		) FROM files f WHERE f.id = c1.featured_id)
		) FROM files f WHERE f.id = c1.featured_id),
		'banner',      (SELECT JSON_BUILD_OBJECT(
		'id', f.id,
		'attributes', (SELECT JSON_BUILD_OBJECT(
			'name',            f.name,
			'alt',             f.alt,
			'caption',         f.caption,
			'ext',             f.ext,
			'provider',        f.provider,
			'width',           f.width,
			'height',          f.height,
			'size',            f.size,
			'url',             f.url
		) FROM files f WHERE f.id = c1.banner_id)
		) FROM files f WHERE f.id = c1.banner_id),
		'collectibles', (SELECT (SELECT JSON_AGG(JSON_BUILD_OBJECT(
		'id', coll.id,
		'attributes', (SELECT JSON_BUILD_OBJECT(
			'item',         collectibles.item_id,
			'tokens',       collectibles.token_ids,
			'collection',   collectibles.collection_id,
			'uri',          collectibles.uri,
			'owner',        collectibles.owner_id,
			'creator',      collectibles.creator_id,
			'details',      collectibles.details,
			'metadata',     collectibles.metadata,
			'created',      collectibles.created
		) FROM collectibles WHERE collectibles.id = coll.id)
		))) FROM collectibles coll WHERE coll.collection_id = c1.id)
	) FROM collections c1 WHERE c1.id = col.id)))) FROM collections col WHERE col.category_id = c.id) as collections
FROM categories c
WHERE c.provider = 'collectible' AND c.parent_id != 0 AND c.enabled = true AND c.deleted isNULL GROUP BY c.id ORDER BY c.order_by ASC;`)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var category Category
		var image AggregatedImageJSON
		collections := make([]*models.MarketplaceCollection, 0)
		err = rows.Scan(&category.ID, &category.Title, &category.Slug, &category.Description, &category.Content, &category.Created, &image, &category.Collections)
		if err != nil {
			return nil, err
		}
		for _, collection := range category.Collections {
			owner, err := users.GetPublicUserByID(ctx, collection.Attributes.Owner)
			if err != nil {
				return nil, err
			}

			creator, err := users.GetPublicUserByID(ctx, collection.Attributes.Creator)
			if err != nil {
				return nil, err
			}

			collectibles := make([]*models.MarketplaceCollectible, 0)

			for _, collectible := range collection.Attributes.Collectibles {
				collectibleOwner, err := users.GetPublicUserByID(ctx, collectible.Attributes.Owner)

				if err != nil {
					return nil, err
				}

				collectibleCreator, err := users.GetPublicUserByID(ctx, collectible.Attributes.Creator)

				if err != nil {
					return nil, err
				}

				collectibles = append(collectibles, &models.MarketplaceCollectible{
					Attributes: &models.MarketplaceCollectibleAttributes{
						ItemID:       collectible.Attributes.Item,
						CollectionID: collectible.Attributes.Collection,
						TokenIds:     collectible.Attributes.Tokens,
						URI:          collectible.Attributes.URI,
						Creator:      collectibleCreator,
						Owner:        collectibleOwner,
						Metadata: &models.MarketplaceCollectibleMetadata{
							Name:        collectible.Attributes.Metadata.Name,
							Description: collectible.Attributes.Metadata.Description,
							Image:       collectible.Attributes.Metadata.Image,
							ExternalURL: collectible.Attributes.Metadata.ExternalURL,
						},
						Details: &models.MarketplaceCollectibleDetails{
							Address:         collectible.Attributes.Details.Address,
							Auction:         collectible.Attributes.Details.Auction,
							Cancelled:       collectible.Attributes.Details.Cancelled,
							Collection:      collectible.Attributes.Details.Collection,
							Fee:             collectible.Attributes.Details.Fee,
							FeeWei:          collectible.Attributes.Details.FeeWei,
							Fulfilled:       collectible.Attributes.Details.Fulfilled,
							Price:           collectible.Attributes.Details.Price,
							PriceWei:        collectible.Attributes.Details.PriceWei,
							Tags:            collectible.Attributes.Details.Tags,
							Total:           collectible.Attributes.Details.Total,
							TotalWei:        collectible.Attributes.Details.TotalWei,
							IsStarted:       collectible.Attributes.Details.IsStarted,
							StartTime:       collectible.Attributes.Details.StartTime,
							EndTime:         collectible.Attributes.Details.EndTime,
							StartPrice:      collectible.Attributes.Details.StartPrice,
							ReservePrice:    collectible.Attributes.Details.ReservePrice,
							StartPriceWei:   collectible.Attributes.Details.StartPriceWei,
							ReservePriceWei: collectible.Attributes.Details.ReservePriceWei,
							Quantity:        collectible.Attributes.Details.Quantity,
						},
					},
					ID: collectible.ID,
				})
			}
			collections = append(collections, &models.MarketplaceCollection{
				ID: collection.ID,
				Attributes: &models.MarketplaceCollectionAttributes{
					Address: common.HexToAddress(collection.Attributes.Address),
					Banner: &models.PublicFile{
						Attributes: collection.Attributes.Banner.Attributes,
						ID:         collection.Attributes.Banner.ID,
					},
					Featured: &models.PublicFile{
						Attributes: collection.Attributes.Featured.Attributes,
						ID:         collection.Attributes.Featured.ID,
					},
					Logo: &models.PublicFile{
						Attributes: collection.Attributes.Logo.Attributes,
						ID:         collection.Attributes.Logo.ID,
					},
					CategoryID:   category.ID,
					Name:         collection.Attributes.Name,
					Symbol:       collection.Attributes.Symbol,
					Description:  collection.Attributes.Description,
					Slug:         collection.Attributes.Slug,
					URL:          collection.Attributes.URL,
					IsApproved:   collection.Attributes.isApproved,
					IsLocked:     collection.Attributes.isLocked,
					Fee:          collection.Attributes.Fee,
					MaxItems:     collection.Attributes.MaxItems,
					Creator:      creator,
					Owner:        owner,
					Created:      strings.Replace(collection.Attributes.Created, "    ", "", 1),
					Collectibles: collectibles,
				},
			})
		}
		if err != nil {
			return nil, err
		}

		publicCategory := &models.PublicCategory{
			Attributes: &models.PublicCategoryAttributes{
				Image: &models.PublicFile{
					Attributes: image.Attributes,
					ID:         image.ID,
				},
				Slug:        category.Slug,
				Title:       category.Title,
				Collections: collections,
			},
			ID: category.ID,
		}

		if category.Content != nil {
			publicCategory.Attributes.Content = *category.Content
		}

		if category.Description != nil {
			publicCategory.Attributes.Description = *category.Description
		}

		categories = append(categories, publicCategory)
	}
	defer func(rows *sqlx.Rows) {
		err = rows.Close()
		if err != nil {

		}
	}(rows)

	return categories, multierr.ErrorOrNil()
}
