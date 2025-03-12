package service

import (
	"go.challenge/models/domain"
	"go.challenge/models/dto"
	"go.challenge/repository"
)

type FavouriteService struct {
	repo *repository.FavouriteRepository
}

func NewFavouriteService(repo *repository.FavouriteRepository) *FavouriteService {
	return &FavouriteService{repo: repo}
}

func (s *FavouriteService) GetFavouritesByUser(userId int64) ([]dto.FavouriteDto, error) {
	items, err := s.repo.FindFavouritesByUser(userId)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *FavouriteService) AddToFavourites(assetId int64, userId int64) {
	newFavourite := domain.Favourite{
		UserId: userId,
		AssetId: assetId,
	}
	s.repo.AddToFavourites(&newFavourite)
}

func (s *FavouriteService) RemoveFromFavourites(favId int64) error {
	return s.repo.RemoveFromFavourites(favId)
}
