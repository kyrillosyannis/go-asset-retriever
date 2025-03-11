package repository

import (
	"go.challenge/models/domain"
	"gorm.io/gorm"
)

type AssetRepository struct {
	db *gorm.DB
}

func NewAssetRepository(db *gorm.DB) *AssetRepository {
	return &AssetRepository{db: db}
}

func (r *AssetRepository) UpdateDescription(id int64, description string) (*domain.Asset, error) {
	asset := &domain.Asset{Id: id}
	if err := r.db.Model(&asset).Update("description", description).Error; err != nil {
		return nil, err
	}
	return asset, nil
}