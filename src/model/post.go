package model

import "time"

type Post struct {
	ID         uint16    `gorm:"primary_key;type:mediumint(8) NOT NULL AUTO_INCREMENT;"`
	UserId     uint16    `gorm:"type:mediumint(11) unsigned not null;default:0;"`
	Author     string    `gorm:"type:varchar(15) NOT NULL;DEFAULT:'';"`
	Title      string    `gorm:"type:varchar(100) NOT NULL;DEFAULT:'';"`
	Color      string    `gorm:"type:varchar(7) NOT NULL;DEFAULT:'';"`
	UrlName    string    `gorm:"type:varchar(100) NOT NULL;DEFAULT:'';"`
	UrlType    int8      `gorm:"type:tinyint(3) not null;DEFAULT:0"`
	Content    string    `gorm:"type:mediumtext NOT NULL"`
	Tags       int8      `gorm:"type:varchar(100) NOT NULL;DEFAULT '';"`
	Views      string    `gorm:"type:mediumint(11) unsigned not null;"`
	Status     int8      `gorm:"type:tinyint(3) NOT NULL;DEFAULT 0;"`
	PostTime   time.Time `gorm:"type:timestamp not null;default:current_timestamp;"`
	UpdateTime time.Time `gorm:"type:timestamp on update current_timestamp;default:current_timestamp;"`
	IsTop      int8      `gorm:"type:tinyint(3) NOT NULL;DEFAULT 0;"`
}
