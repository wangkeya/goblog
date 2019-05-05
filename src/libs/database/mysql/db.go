package mysql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open("mysql", "root:111111@/goblog?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	err = DB.DB().Ping()
	if err != nil {
		panic(err)
	}

	//设置默认表名前缀
	gorm.DefaultTableNameHandler = func(DB *gorm.DB, defaultTableName string) string {
		return "t_" + defaultTableName
	}
	DB.SingularTable(true)
	//连接池
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
}

func CloseDb()  {
	DB.Close()
}
