package model

type User struct {
	Id       int    `gorm:"AUTO_INCREMENT" form:"id"`
	Username string `gorm:"not null" form:"username"`
	Phone    string `gorm:"not null" form:"phone"`
}
