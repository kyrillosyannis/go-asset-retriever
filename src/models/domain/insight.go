package domain

type Insight struct {
	Id int64 `gorm:"primaryKey"`
	Text string `gorm:"column:text"`
	AssetId int64 `gorm:"column:asset_id;uniqueIndex"`
	Asset Asset `gorm:"foreignKey:AssetId"`
}
