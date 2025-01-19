package repository

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/ole-larsen/plutonium/models"
)

type AggregatedMenuAttributesJSON models.PublicMenuAttributes

func (a *AggregatedMenuAttributesJSON) Scan(val interface{}) error {
	switch v := val.(type) {
	case []byte:
		_ = json.Unmarshal(v, &a)
		return nil
	case string:
		_ = json.Unmarshal([]byte(v), &a)
		return nil
	default:
		return nil
	}
}

func (a *AggregatedMenuAttributesJSON) Value() (driver.Value, error) {
	return json.Marshal(a)
}

type AggregatedSliderItemJSON []*models.PublicSliderItem

func (a *AggregatedSliderItemJSON) Scan(val interface{}) error {
	switch v := val.(type) {
	case []byte:
		_ = json.Unmarshal(v, &a)
		return nil
	case string:
		_ = json.Unmarshal([]byte(v), &a)
		return nil
	default:
		return nil
	}
}

func (a *AggregatedSliderItemJSON) Value() (driver.Value, error) {
	return json.Marshal(a)
}

type AggregatedCategoryCollectionsCollectiblesJSON []CategoryCollectionCollectible

func (a *AggregatedCategoryCollectionsCollectiblesJSON) Scan(val interface{}) error {
	switch v := val.(type) {
	case []byte:
		_ = json.Unmarshal(v, &a)
		return nil
	case string:
		_ = json.Unmarshal([]byte(v), &a)
		return nil
	default:
		return nil
	}
}

func (a *AggregatedCategoryCollectionsCollectiblesJSON) Value() (driver.Value, error) {
	return json.Marshal(a)
}

type AggregatedCategoryCollectionsJSON []CategoryCollection

func (a *AggregatedCategoryCollectionsJSON) Scan(val interface{}) error {
	switch v := val.(type) {
	case []byte:
		_ = json.Unmarshal(v, &a)
		return nil
	case string:
		_ = json.Unmarshal([]byte(v), &a)
		return nil
	default:
		return nil
	}
}

func (a *AggregatedCategoryCollectionsJSON) Value() (driver.Value, error) {
	return json.Marshal(a)
}

type AggregatedImageJSON models.PublicFile

func (a *AggregatedImageJSON) Scan(val interface{}) error {
	switch v := val.(type) {
	case []byte:
		_ = json.Unmarshal(v, &a)
		return nil
	case string:
		_ = json.Unmarshal([]byte(v), &a)
		return nil
	default:
		return nil
	}
}

func (a *AggregatedImageJSON) Value() (driver.Value, error) {
	return json.Marshal(a)
}

type AggregatedWallet []*models.Wallet

func (a *AggregatedWallet) Scan(val interface{}) error {
	switch v := val.(type) {
	case []byte:
		_ = json.Unmarshal(v, &a)
		return nil
	case string:
		_ = json.Unmarshal([]byte(v), &a)
		return nil
	default:
		return nil
	}
}

func (a *AggregatedWallet) Value() (driver.Value, error) {
	return json.Marshal(a)
}

type AggregatedSocial []*models.Social

func (a *AggregatedSocial) Scan(val interface{}) error {
	switch v := val.(type) {
	case []byte:
		_ = json.Unmarshal(v, &a)
		return nil
	case string:
		_ = json.Unmarshal([]byte(v), &a)
		return nil
	default:
		return nil
	}
}

func (a *AggregatedSocial) Value() (driver.Value, error) {
	return json.Marshal(a)
}

type AggregatedAuthorJSON models.PublicAuthor

func (a *AggregatedAuthorJSON) Scan(val interface{}) error {
	switch v := val.(type) {
	case []byte:
		_ = json.Unmarshal(v, &a)
		return nil
	case string:
		_ = json.Unmarshal([]byte(v), &a)
		return nil
	default:
		return nil
	}
}

func (a *AggregatedAuthorJSON) Value() (driver.Value, error) {
	return json.Marshal(a)
}

type AggregatedAuthorAttributes models.PublicAuthorAttributes

func (a *AggregatedAuthorAttributes) Scan(val interface{}) error {
	switch v := val.(type) {
	case []byte:
		_ = json.Unmarshal(v, &a)
		return nil
	case string:
		_ = json.Unmarshal([]byte(v), &a)
		return nil
	default:
		return nil
	}
}

func (a *AggregatedAuthorAttributes) Value() (driver.Value, error) {
	return json.Marshal(a)
}

type AggregatedPageAttributes models.PublicPageAttributes

func (a *AggregatedPageAttributes) Scan(val interface{}) error {
	switch v := val.(type) {
	case []byte:
		_ = json.Unmarshal(v, &a)
		return nil
	case string:
		_ = json.Unmarshal([]byte(v), &a)
		return nil
	default:
		return nil
	}
}

func (a *AggregatedPageAttributes) Value() (driver.Value, error) {
	return json.Marshal(a)
}

type AggregatedContactAttributes models.PublicContactAttributes

func (a *AggregatedContactAttributes) Scan(val interface{}) error {
	switch v := val.(type) {
	case []byte:
		_ = json.Unmarshal(v, &a)
		return nil
	case string:
		_ = json.Unmarshal([]byte(v), &a)
		return nil
	default:
		return nil
	}
}

func (a *AggregatedContactAttributes) Value() (driver.Value, error) {
	return json.Marshal(a)
}

type AggregatedTagJSON []models.PublicTag

func (a *AggregatedTagJSON) Scan(val interface{}) error {
	switch v := val.(type) {
	case []byte:
		_ = json.Unmarshal(v, &a)
		return nil
	case string:
		_ = json.Unmarshal([]byte(v), &a)
		return nil
	default:
		return nil
	}
}

func (a *AggregatedTagJSON) Value() (driver.Value, error) {
	return json.Marshal(a)
}
