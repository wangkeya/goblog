package main

import (
	"github.com/gin-gonic/gin"
	"router"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	_ "common"
	"io"
)

const ROOT = "/Users/alan/goblog"

func init()  {
	//save pid
	pidFilename := ROOT + "/pid/" + filepath.Base(os.Args[0]) + ".pid"
	pid := os.Getpid()
	ioutil.WriteFile(pidFilename, []byte(strconv.Itoa(pid)), 0755)
}

func main() {
	r := gin.New()
	// 设置日志文件
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	// 使用日志中间件
	r.Use(gin.Logger())
	// 设置静态文件夹
	r.Static("/static", "./static")
	// 加载路由
	router.LoadWebRouters(r)
	router.LoadAdminRouters(r)
	r.Run(":8080")
}
