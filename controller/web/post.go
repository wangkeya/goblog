package web

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"goblog/service"
	"net/http"
	"fmt"
	"goblog/libs"
)

type PostController struct{}

func (p *PostController) GetInfo(c *gin.Context) {
	var postId string
	postId = c.Param("id")
	id, _ := strconv.Atoi(postId)
	postService := service.PostService{}
	post, _ := postService.GetPostById(id)
	fmt.Println(libs.FormatAsDate(post.UpdateTime))

	// 返回一个json格式的数据
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": post,
	})
}

func (p *PostController) GetList(c *gin.Context)  {

	page,_ := strconv.Atoi(c.Query("page"))
	pageRows,_ := strconv.Atoi(c.Query("pageRows"))
	var data = make(map[string]interface{})

	postService :=service.PostService{}
	posts,count,_ := postService.PageData(page,pageRows)
	data["lists"] = posts
	data["total"] = count

	// 返回一个json格式的数据
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": data,
	})
}

func (p *PostController) Edit(c *gin.Context)  {
	postService :=service.PostService{}
	postService.Insert()

	// 返回一个json格式的数据
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": "",
	})
}

func (p *PostController) Update(c *gin.Context)  {
	postService :=service.PostService{}
	postService.Update()
	// 返回一个json格式的数据
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": "",
	})
}

func (p *PostController) Delete(c *gin.Context)  {
	id := c.Param("id")
	_id,_ := strconv.Atoi(id)
	postService := service.PostService{}
	err := postService.Delete(_id)
	fmt.Println(err)

	// 返回一个json格式的数据
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "删除成功",
		"data": id,
	})
}

