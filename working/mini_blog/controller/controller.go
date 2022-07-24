package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginGet(ctx *gin.Context) {
	username := ctx.Query("username")
	password := ctx.Query("password")
	ctx.String(200, "username=%v password=%v", username, password)

	//ctx.HTML(http.StatusOK, "login.html", gin.H{"title": "登录"}) //登录页面模板渲染

}

// Login 登录
func LoginPost(ctx *gin.Context) {

	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	ctx.JSON(200, gin.H{ //返回json数据类型
		"username": username,
		"password": password,
	})

	//ctx.Redirect(http.StatusMovedPermanently, "register.html")  //已有账号 输入账号 点击登录 跳转到blog首页
}

// Register 注册
func Register(ctx *gin.Context) {
	ctx.Redirect(http.StatusMovedPermanently, "login.html")        //注册完成 点击注册按钮 跳转到登录页面
	ctx.HTML(http.StatusOK, "register.html", gin.H{"title": "注册"}) //注册页面模板渲染
}

/*func LoginPost(c *gin.Context) {
	var res = gin.H{}
	// 获取表单数据
	username := c.PostForm("username")
	password := c.PostForm("password")

	user := &models.UserLoginMessage{
		UserName: username,
		Password: password,
	}
	if len(user.UserName) == 0 || len(user.Password) == 0 {
		// 另一种json格式的返回
		res["message"] = "用户名或密码不能为空！"
		return
	}
}*/
