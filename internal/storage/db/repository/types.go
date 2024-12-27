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
