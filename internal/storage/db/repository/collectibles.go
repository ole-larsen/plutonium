package repository

import (
	"time"
)

type CollectibleMetadata struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	ExternalURL string `json:"external_url"`
}

type CollectibleDetails struct {
	Address         string `json:"address"`
	Tags            string `json:"tags"`
	Collection      string `json:"collection"`
	Auction         bool   `json:"auction"`
	Fulfilled       bool   `json:"fulfilled"`
	Cancelled       bool   `json:"cancelled"`
	PriceWei        string `json:"price_wei"`
	Price           string `json:"price"`
	TotalWei        string `json:"total_wei"`
	Total           string `json:"total"`
	FeeWei          string `json:"fee_wei"`
	Fee             string `json:"fee"`
	IsStarted       bool   `json:"is_started"`
	StartTime       int64  `json:"start_time"`
	EndTime         int64  `json:"end_time"`
	StartPrice      string `json:"start_price"`
	ReservePrice    string `json:"reserve_price"`
	StartPriceWei   string `json:"start_price_wei"`
	ReservePriceWei string `json:"reserve_price_wei"`
	Quantity        int64  `json:"quantity"`
}

type Collectible struct {
	ID           int64               `db:"id"`
	TokenIDs     []int64             `db:"token_ids"`
	CollectionID int64               `db:"collection_id"`
	Creator      string              `db:"creator"`
	Owner        string              `db:"owner"`
	URI          string              `db:"uri"`
	Details      CollectibleDetails  `db:"details"`
	Metadata     CollectibleMetadata `db:"metadata"`
	Created      time.Time           `db:"created"`
	Updated      time.Time           `db:"updated"`
	Deleted      time.Time           `db:"deleted"`
}

// type CollectibleRepository interface {
// 	Create(collectibleMap map[string]interface{}, details CollectibleDetails, metadata models.Metadata) error
// 	Buy(collectibleMap map[string]interface{}, details CollectibleDetails) error
// 	Sell(collectibleMap map[string]interface{}, details CollectibleDetails) error
// }

// type CollectibleRepo struct {
// 	tbl string
// 	db  *sqlx.DB
// 	ctx context.Context
// }

// func NewCollectibleRepo(ctx context.Context, store db.Storer) *CollectibleRepo {
// 	return &CollectibleRepo{
// 		tbl: "collectibles",
// 		db:  store.InnerDB(),
// 		ctx: ctx,
// 	}
// }

// func (c *CollectibleRepo) Create(collectibleMap map[string]interface{}, details CollectibleDetails, metadata models.Metadata) error {
// 	if c.db == nil {
// 		return errDbNotInitialized
// 	}
// 	detailsJSON, err := json.Marshal(details)
// 	if err != nil {
// 		return err
// 	}

// 	collectibleMap["details"] = string(detailsJSON)

// 	metadataJSON, err := json.Marshal(metadata)
// 	if err != nil {
// 		return err
// 	}

// 	collectibleMap["metadata"] = string(metadataJSON)

// 	_, err = c.db.NamedExec(`
// 		INSERT INTO collectibles (item_id, token_ids, collection_id, creator_id, owner_id, uri, details, metadata)
// 		VALUES (:item_id, :token_ids, :collection_id, :creator_id, :owner_id, :uri, :details, :metadata)
// 		ON CONFLICT
// 		DO
// 	    NOTHING;`, collectibleMap)
// 	return err
// }

// func (c *CollectibleRepo) Buy(collectibleMap map[string]interface{}, details CollectibleDetails) error {
// 	if c.db == nil {
// 		return errDbNotInitialized
// 	}
// 	detailsJSON, err := json.Marshal(details)
// 	if err != nil {
// 		return err
// 	}

// 	collectibleMap["details"] = string(detailsJSON)
// 	_, err = c.db.NamedExec(`
// 	UPDATE collectibles SET
// 		details=:details,
// 		owner_id=:owner_id
// 	WHERE token_id=:token_id AND collection_id=:collection_id`, collectibleMap)
// 	if err != nil {
// 		return err
// 	}
// 	return err
// }

// func (c *CollectibleRepo) Sell(collectibleMap map[string]interface{}, details CollectibleDetails) error {
// 	if c.db == nil {
// 		return errDbNotInitialized
// 	}
// 	detailsJSON, err := json.Marshal(details)
// 	if err != nil {
// 		return err
// 	}

// 	collectibleMap["details"] = string(detailsJSON)
// 	_, err = c.db.NamedExec(`
// 	UPDATE collectibles SET
// 		details=:details,
// 		owner_id=:owner_id
// 	WHERE token_id=:token_id AND collection_id=:collection_id`, collectibleMap)
// 	if err != nil {
// 		return err
// 	}
// 	return err
// }
