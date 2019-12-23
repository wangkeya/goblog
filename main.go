package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"goblog/middleware"
	"os"
	"goblog/router"
	"goblog/libs"
	"log"
	"time"
)

var (
	ProjectPath = libs.GetCurrentPath()
)

func main() {
	r := gin.New()
	// 设置日志文件
	logPath := libs.CreateDir(ProjectPath+"/runtime/log",false)
	f, err := os.Create( logPath + "/gin-"+libs.FormatAsDate(time.Now())+".log")
	if err != nil {
		log.Fatal(err)
	}
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	// 使用日志中间件
	r.Use(gin.Logger(),gin.Recovery())
	// 设置静态文件夹
	r.Static("/static", "./static")
	//跨域
	r.Use(middleware.CrossDomain())
	// 加载路由
	router.LoadWebRouters(r)
	router.LoadAdminRouters(r)
	r.Run(":5200")
}
