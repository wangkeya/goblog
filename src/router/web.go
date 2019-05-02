package router

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// LoadRouters 初始化router
func LoadWebRouters(router *gin.Engine) {
	loadWebRouters(router)
}

func loadWebRouters(router *gin.Engine) {
	// 这里测试根路由
	/*
		控制器函数会接收一个`gin.Context`类型的指针，这个指针包含http的请求和响应信息和操作方法。
	*/
	router.GET("/api", func(c *gin.Context) {
		// 返回一个json格式的数据
		c.JSON(http.StatusOK, gin.H{
			"Status": 0,
			"data": "ok",
		})
	})
}