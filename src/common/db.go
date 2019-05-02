package common

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)


func init()  {
	db, err := sql.Open("mysql", "root:111111@/goblog")

	if err != nil{
		panic(err.Error())
		log.Println(err)
		return
	}

	defer db.Close()

}