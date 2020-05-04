package web

import (
	"github.com/gin-gonic/gin"
	"goblog/service"
	"net/http"
	"strconv"
)

type TagController struct{}

func (tagController *TagController) GetTagById(c *gin.Context) {
	var tagId string
	tagId = c.Query("tagId")
	id, _ := strconv.Atoi(tagId)
	tagService := service.TagService{}
	tag, _ := tagService.GetTagById(id)

	// 返回一个json格式的数据
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": tag,
	})
}

