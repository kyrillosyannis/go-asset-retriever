package service

import (
	"go.challenge/models/domain"
	"go.challenge/repository"
)

type AssetService struct {
	repo *repository.AssetRepository
}

func NewAssetService(repo *repository.AssetRepository) *AssetService {
	return &AssetService{repo: repo}
}

func (s *AssetService) UpdateDescription(assetId int64, description string) (*domain.Asset, error) {
	return s.repo.UpdateDescription(assetId, description)
}
