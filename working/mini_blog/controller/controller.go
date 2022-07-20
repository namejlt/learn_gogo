package controller

import "github.com/gin-gonic/gin"

func Welcome(ctx *gin.Context) {
	ctx.String(200, "Welcome to mini_blog!")
}

func Login(ctx *gin.Context) {
	ctx.HTML(200, "login.html", nil) //登录页面模板渲染
}
func Register(ctx *gin.Context) {

	ctx.HTML(200, "register.html", nil) //注册页面模板渲染
}
