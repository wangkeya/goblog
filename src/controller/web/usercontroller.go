package web

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	
} 

func (user *UserController) GetUsers(c *gin.Context) {

	// 返回一个json格式的数据
	c.JSON(http.StatusOK, gin.H{
		"Status": 0,
		"data": "ok",
	})
}

