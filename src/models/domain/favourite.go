package domain

type Favourite struct {
	Id int64 `gorm:"primaryKey"`
	UserId int64 `gorm:"column:user_id"`
	AssetId int64 `gorm:"column:asset_id"`
	Asset Asset `gorm:"foreignKey:AssetID"`
}
