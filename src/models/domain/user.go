package domain

type User struct {
	Id int64 `gorm:"primaryKey`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
}
