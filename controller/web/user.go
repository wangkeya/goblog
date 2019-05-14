package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct{}

func (user *UserController) GetUsers(c *gin.Context) {

	// 返回一个json格式的数据
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": "",
	})
}

