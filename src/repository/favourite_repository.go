package repository

import (
	"go.challenge/models/dto"
	"gorm.io/gorm"
)

type FavouriteRepository struct {
	db *gorm.DB
}

func NewFavouriteRepository(db *gorm.DB) *FavouriteRepository {
	return &FavouriteRepository{db: db}
}

func (r *FavouriteRepository) FindFavouritesByUser(userId int64) ([]dto.FavouriteDto, error) {
	var favourites []dto.FavouriteDto
	err := r.db.Table("favourites f").
		Select(`
			f.id AS favourite_id,
			a.id AS asset_id,
			a.description,
			c.title AS chart_title,
			c.x_title AS x_axis_label,
			c.y_title AS y_axis_label,
			i.text AS insight_text,
			au.gender,
			au.country_of_birth,
			au.age_group,
			au.daily_hours_on_social_media,
			au.last_month_number_of_purchases
		`).
		Joins("JOIN assets a ON f.asset_id = a.id").
		Joins("LEFT JOIN charts c ON a.id = c.asset_id").
		Joins("LEFT JOIN insights i ON a.id = i.asset_id").
		Joins("LEFT JOIN audiences au ON a.id = au.asset_id").
		Where("f.user_id = ?", userId).
		Scan(&favourites).Error
		return favourites, err
}
