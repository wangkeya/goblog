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
	mysql "goblog/libs/database/mysql"
	redis "goblog/libs/redis"
	"goblog/libs"
)

var (
	ProjectPath string = libs.GetCurrentPath()
)

func init() {
	//save pid
	pidFilename :=  ProjectPath + "/runtime/pid/" + filepath.Base(os.Args[0]) + ".pid"
	pid := os.Getpid()
	ioutil.WriteFile(pidFilename, []byte(strconv.Itoa(pid)), 0755)
}

func main() {
	defer mysql.CloseDb()
	err := redis.RedisClient.Set("test", "test", 0).Err()
	if err != nil {
		panic(err)
	}

	r := gin.New()
	// 设置日志文件
	f, _ := os.Create( ProjectPath + "/runtime/log/gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	// 使用日志中间件
	r.Use(gin.Logger())
	// 设置静态文件夹
	r.Static("/static", "./static")
	//跨域
	r.Use(middleware.CrossDomain())
	// 加载路由
	router.LoadWebRouters(r)
	router.LoadAdminRouters(r)

	r.Run(":5200")
}
