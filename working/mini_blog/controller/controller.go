package controller

import "github.com/gin-gonic/gin"

func Welcome(ctx *gin.Context) {
	ctx.String(200, "Welcome to mini_blog!")
}

func Login(ctx *gin.Context) {
	ctx.HTML(200, "login.html", nil)
}
func Register(ctx *gin.Context) {

	ctx.HTML(200, "register.html", nil)
}
