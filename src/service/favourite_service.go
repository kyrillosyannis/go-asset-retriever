package service

import (
	"fmt"
	"go.challenge/models/dto"
	"go.challenge/repository"
)

type FavouriteService struct {
	repo *repository.FavouriteRepository
}

func NewFavouriteService(repo *repository.FavouriteRepository) *FavouriteService {
	return &FavouriteService{repo: repo}
}

func (s *FavouriteService) GetFavouritesByUser() ([]dto.FavouriteDto, error) {
	items, err := s.repo.FindFavouritesByUser(1)
	if err != nil {
		return nil, err
	}

    fmt.Println(items)

	return items, nil
}
