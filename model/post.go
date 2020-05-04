package model

import "time"

type Post struct {
	ID         int       `gorm:"primary_key;type:mediumint(8) NOT NULL AUTO_INCREMENT;" json:"id"`
	UserId     uint16    `gorm:"type:mediumint(11) unsigned not null;default:0;" json:"user_id"`
	Author     string    `gorm:"type:varchar(15) NOT NULL;DEFAULT:'';" json:"author"`
	Title      string    `gorm:"type:varchar(100) NOT NULL;DEFAULT:'';" json:"title"`
	Color      string    `gorm:"type:varchar(7) NOT NULL;DEFAULT:'';" json:"color"`
	UrlName    string    `gorm:"type:varchar(100) NOT NULL;DEFAULT:'';" json:"url_name"`
	UrlType    int8      `gorm:"type:tinyint(3) not null;DEFAULT:0" json:"url_type"`
	Content    string    `gorm:"type:mediumtext NOT NULL" json:"content"`
	Tags       string    `gorm:"type:varchar(100) NOT NULL;DEFAULT '';" json:"tags"`
	Views      string    `gorm:"type:mediumint(11) unsigned not null;" json:"views"`
	Status     int8      `gorm:"type:tinyint(3) NOT NULL;DEFAULT 0;" json:"status"`
	PostTime   time.Time `gorm:"type:timestamp not null;default:current_timestamp;" json:"post_time"`
	UpdateTime time.Time `gorm:"type:timestamp on update current_timestamp;default:current_timestamp;" json:"update_time"`
	IsTop      int8      `gorm:"type:tinyint(3) NOT NULL;DEFAULT 0;" json:"is_top"`
}
