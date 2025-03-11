package domain

type Audience struct {
	Id int64 `gorm:"primaryKey"`
	Gender string `gorm:"column:gender"`
	CountryOfBirth string `gorm:"column:country_of_birth"`
	AgeGroup string `gorm:"column:age_group"`
	DailyHoursOnSocialMedia int16 `gorm:"column:daily_hours_on_social_media"`
	LastMonthNumberOfPurchases int `gorm:"column:last_month_number_of_purchases"`
	AssetId int64 `gorm:"column:asset_id;uniqueIndex"`
	Asset Asset `gorm:"foreignKey:AssetId"`
}
