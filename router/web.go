package router

import (
	ctrls "goblog/controller/web"
	"github.com/gin-gonic/gin"
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
		v1.GET("/user/getUsers", userController.GetUsers)

		postController := new(ctrls.PostController)
		v1.GET("/post/getInfo/:id", postController.GetInfo)
		v1.GET("/post/getList",postController.GetList)
		v1.POST("/post/edit",postController.Edit)
		v1.POST("/post/update",postController.Update)
		v1.GET("/post/delete/:id",postController.Delete)

		tagController := new(ctrls.TagController)
		v1.GET("/tag/getTagById", tagController.GetTagById)
	}

}
