package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"mini_blog/routers"
)

func main() {

	//默认路由引擎
	router := gin.Default()

	//前端页面模板解析
	router.LoadHTMLGlob("views/*")

	//路由分离
	routers.ParentRouter(router)

	//启动http服务 默认端口号为127.0.0.1:8080
	router.Run(":8080")
}
