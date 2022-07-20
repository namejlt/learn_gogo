package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"mini_blog/routers"
)

func main() {
	//用户名：密码@tcp（ip:port）/数据库？charset=utf8&parseTime=true&loc=Local
	db, err := gorm.Open("mysql", "root:kk123456yy@tcp(localhost:3306)/weibo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	defer db.Close() //关闭空闲连接，防止内存泄漏

	//默认路由引擎
	router := gin.Default()

	//前端页面模板解析
	router.LoadHTMLGlob("views/*")

	//路由分离
	routers.ParentRouter(router)

	//启动http服务 默认端口号为127.0.0.1:8080
	router.Run()
}
