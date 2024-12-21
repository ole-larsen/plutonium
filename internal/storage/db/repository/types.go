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
