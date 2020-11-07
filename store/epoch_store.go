package store

import (
	"gorm.io/gorm"

	"github.com/figment-networks/filecoin-indexer/model"
)

// EpochStore handles database operations on epochs
type EpochStore struct {
	db *gorm.DB
}

// Create stores an epoch record
func (es *EpochStore) Create(epoch *model.Epoch) error {
	return es.db.Create(epoch).Error
}

// LastHeight returns the most recent height
func (es *EpochStore) LastHeight() (int64, error) {
	var result int64

	err := es.db.Table("epochs").Select("MAX(height)").Scan(&result).Error
	if err != nil {
		return 0, err
	}

	return result, nil
}
