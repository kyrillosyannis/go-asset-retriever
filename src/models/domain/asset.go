package domain

type Asset struct {
	Id int64 `gorm:"primaryKey"`
	Description string `gorm:"column:description"`
}
