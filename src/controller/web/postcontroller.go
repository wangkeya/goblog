package web

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"service"
	"net/http"
)

type PostController struct{}

func (postController *PostController) GetPostById(c *gin.Context) {
	var postId string
	postId = c.Query("postId")
	id, _ := strconv.Atoi(postId)
	postService := service.PostService{}
	post, _ := postService.GetPostById(id)

	// 返回一个json格式的数据
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": post,
	})
}

