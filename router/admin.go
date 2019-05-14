package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	// ctrs "github.com/zachrey/blog/controllers"
)

// LoadRouters 初始化router
func LoadAdminRouters(router *gin.Engine) {
	loadAdminRouters(router)
}

func loadAdminRouters(router *gin.Engine) {
	/*
		控制器函数会接收一个`gin.Context`类型的指针，这个指针包含http的请求和响应信息和操作方法。
	*/
	router.GET("/", func(c *gin.Context) {
		// 返回一个json格式的数据
		c.JSON(http.StatusOK, gin.H{
			"Status": 0,
			"data":   "ok",
		})
	})
}
