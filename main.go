package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"goblog/middleware"
	"os"
	"path/filepath"
	"goblog/router"
	"strconv"
	"goblog/libs"
	"log"
)

var (
	ProjectPath = libs.GetCurrentPath()
)

func init() {
	//save pid
	libs.CreateDir(ProjectPath+"/runtime/pid",false)
	pidFilename :=  ProjectPath + "/runtime/pid/" + filepath.Base(os.Args[0]) + ".pid"
	pid := os.Getpid()
	ioutil.WriteFile(pidFilename, []byte(strconv.Itoa(pid)), 0755)
}

func main() {
	r := gin.New()
	// 设置日志文件
	logPath := libs.CreateDir(ProjectPath+"/runtime/log",true)
	f, err := os.Create( logPath + "/gin.log")
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
