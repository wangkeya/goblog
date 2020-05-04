package model

type User struct {
	Id       int    `gorm:"AUTO_INCREMENT"  json:"id"`
	Username string `gorm:"not null"  json:"username"`
	Phone    string `gorm:"not null"  json:"phone"`
}
