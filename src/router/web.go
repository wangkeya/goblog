package router

import (
	"github.com/gin-gonic/gin"
	ctrls "controller/web"
)

// LoadRouters 初始化router
func LoadWebRouters(router *gin.Engine) {
	loadWebRouters(router)
}

func loadWebRouters(router *gin.Engine) {
	/*
		控制器函数会接收一个`gin.Context`类型的指针，这个指针包含http的请求和响应信息和操作方法。
	*/
	v1 := router.Group("/web")
	{
		userController := new(ctrls.UserController)
		v1.GET("/getUsers",userController.GetUsers)
	}

}