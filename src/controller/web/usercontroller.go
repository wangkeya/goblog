package web

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"service"
	"fmt"
	"strconv"
)

type UserController struct {}

func (user *UserController) GetUsers(c *gin.Context) {
	var tagId string
	tagId = c.Query("tagId")
	fmt.Println(tagId)
	id,_ := strconv.Atoi(tagId)
	tagService := service.TagService{}
	tag,_ := tagService.GetTagById(id)

	// 返回一个json格式的数据
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":"获取成功",
		"data": tag,
	})
}

