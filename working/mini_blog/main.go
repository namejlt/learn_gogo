package main

import (
	"github.com/gin-gonic/gin"
	"mini_blog/controller"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("views/*")

	router.GET("/", controller.Login)            //初始化接口
	router.GET("/login", controller.Login)       //登录接口
	router.GET("/register", controller.Register) //注册接口

	router.Run()
}
