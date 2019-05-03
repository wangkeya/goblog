package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func NewMysql() (db *gorm.DB) {
	var err error
	db, err = gorm.Open("mysql", "root:111111@/goblog?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	//设置默认表名前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "t_" + defaultTableName
	}
	db.SingularTable(true)

	return
}
