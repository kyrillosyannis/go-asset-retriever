package domain

import "time"

type Chart struct {
	Id           int64 `gorm:"primaryKey"`
	Title        string `gorm:"column:title"`
	XAxisTitle   string `gorm:"column:x_title"`
	YAxisTitle   string	`gorm:"column:y_title"`
	XCoordinate  float64 `gorm:"column:x_coord"`
	YCoordinate  float64 `gorm:"column:y_coord"`
	CreationDate time.Time `gorm:"column:creation_date"`
	UpdateDate 	 time.Time `gorm:"column:update_date"`
	AssetId int64 `gorm:"column:asset_id;uniqueIndex"`
	Asset Asset `gorm:"foreignKey:AssetID"`
}
