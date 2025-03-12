package dto

type FavouriteDto struct {
		FavoriteId int64 `json:"favorite_id"`
		AssetId int64 `json:"asset_id"`
		Description string `json:"description"`
		Type string `json:"type"`
		ChartTitle string `json:"chart_title"`
		XAxisLabel string `json:"x_axis_label"`
		YAxisLabel string `json:"y_axis_label"`
		XCoordinate float64 `json:"x_coordinate"`
		YCoordinate float64 `json:"y_coordinate"`
		InsightText string `json:"insight_text"`
		Gender string `json:"gender"`
		CountryOfBirth string `json:"country_of_birth"`
		AgeGroup string `json:"age_group"`
		DailyHoursOnSocialMedia int16 `json:"daily_hours_on_social_media"`
		LastMonthNumberOfPurchases int `json:"last_month_number_of_purchases"`
	}
