package web

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"service"
)

type UserController struct {
	
} 

func (user *UserController) GetUsers(c *gin.Context) {
	tagService := service.TagService{}
	tag,_ := tagService.GetTagById(1)



	// 返回一个json格式的数据
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":"获取成功",
		"data": tag,
	})
}

