package repository

import (
	"fmt"

	"go.challenge/models/domain"
	"gorm.io/gorm"
)

// Repository layer (similar to Spring Data JPA Repository)
type ChartRepository struct {
	db *gorm.DB
}

func NewChartRepository(db *gorm.DB) *ChartRepository {
	return &ChartRepository{db: db}
}

func (r *ChartRepository) Create(item *domain.Chart) error {
	return r.db.Create(item).Error
}

func (r *ChartRepository) FindAll() ([]domain.Chart, error) {
	var items []domain.Chart
	fmt.Println(r)
	result := r.db.Find(&items)
	return items, result.Error
}

func (r *ChartRepository) FindByID(id uint) (*domain.Chart, error) {
	var item domain.Chart
	result := r.db.First(&item, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &item, nil
}