package config

import (
	"fmt"
	"github.com/go-ini/ini"
	"os"
)

var Config *ini.File
var RootPath string

func init() {
	//RootPath项目绝对路径
	RootPath = "/Users/alan/goblog/src"
	var err error
	Config, err = ini.Load(RootPath + "/conf/app.ini")
	database := Config.Section("database")
	username, _ := database.GetKey("username")
	fmt.Println(username)

	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
}
