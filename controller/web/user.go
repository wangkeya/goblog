package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"goblog/service"
	"log"
	"time"
	jwtgo "github.com/dgrijalva/jwt-go"
	myjwt "goblog/middleware"
	"goblog/model"
)

type UserController struct{}

func (u *UserController) Register(c *gin.Context) {
	// 返回一个json格式的数据
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": "",
	})
}

func (u *UserController) Login(c *gin.Context)()  {

	userService := service.UserService{}
	isPass, user, err := userService.CheckLogin(1)
	if err != nil {
		log.Fatal(err)
	}
	if isPass {
		generateToken(c, user)
	}else {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "验证失败," + err.Error(),
		})
	}
}

// 生成令牌
func generateToken(c *gin.Context, user *model.User) {
	j := &myjwt.JWT{
		[]byte("newtrekWang"),
	}
	claims := myjwt.CustomClaims{
		user.Id,
		user.Username,
		user.Phone,
		jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
			Issuer:    "newtrekWang",                   //签名的发行者
		},
	}

	token, err := j.CreateToken(claims)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    err.Error(),
		})
		return
	}

	log.Println(token)

	// LoginResult 登录结果结构
	type LoginResult struct {
		Token string `json:"token"`
		*model.User
	}
	data := LoginResult{
		User:  user,
		Token: token,
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":    "登录成功！",
		"data":   data,
	})
	return
}

func (u *UserController) GetInfo(c *gin.Context)()  {

	claims,_ := c.Get("claims")

	//userService := service.UserService{}
	//user, err := userService.GetInfo(1)
	//if err != nil {
	//	log.Fatal(err)
	//}

	// 返回一个json格式的数据
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "登录成功",
		"data": &claims,
	})
}