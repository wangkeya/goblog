package model

type Tag struct {
	ID    uint16 `gorm:"primary_key;type:mediumint(8) NOT NULL AUTO_INCREMENT;" json:"id"`
	Name  string `gorm:"type:varchar(20) NOT NULL;DEFAULT:'';" json:"name"`
	Count int8   `gorm:"type:mediumint(11) unsigned not null;default:0;" json:"count"`
}
