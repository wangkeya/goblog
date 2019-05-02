package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	_ "common"
)

const ROOT = "/Users/alan/goblog"

func init()  {
	//save pid
	pidFilename := ROOT + "/pid/" + filepath.Base(os.Args[0]) + ".pid"
	pid := os.Getpid()
	ioutil.WriteFile(pidFilename, []byte(strconv.Itoa(pid)), 0755)
}


func main() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
