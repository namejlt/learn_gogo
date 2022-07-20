package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"mini_blog/models"
	"mini_blog/routers"
)

func main() {
	//用户名：密码@tcp（ip:port）/数据库？charset=utf8&parseTime=true&loc=Local
	db, err := gorm.Open("mysql", "root:kk123456yy@tcp(localhost:3306)/weibo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close() //关闭空闲连接，防止内存泄漏

	//创建三个表  方式一
	if !db.HasTable(&models.User{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8 user '用户基础信息表'").
			CreateTable(&models.User{}).Error; err != nil {
			panic(err)
		}
	}
	if !db.HasTable(&models.UserLoginMessage{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8 userLoginMessage '用户登录信息表'").
			CreateTable(&models.UserLoginMessage{}).Error; err != nil {
			panic(err)
		}
	}
	if !db.HasTable(&models.Message{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8 message '用户发送消息表'").
			CreateTable(&models.Message{}).Error; err != nil {
			panic(err)
		}
	}

	//方式二
	//db.CreateTable(&models.User{}) //表名称为users
	//db.CreateTable(&models.UserLoginMessage{})
	//db.CreateTable(&models.Message{})
	//指定表名称
	//db.Table("user").CreateTable(&models.User{}) //创建表名为user

	//删除表
	//db.DropTable("user") //或者下面方式
	//db.DropTable(&models.User{})

	//检查表是否存在
	/*b := db.HasTable("users")
	fmt.Println(b)
	b2 := db.HasTable(&models.User{})
	fmt.Println(b2)*/

	//方式三  自动迁移表  仅仅会创建表 并添加缺少的列和索引，不会改变现有列的类型或删除未使用的列以保护数据
	/*db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.UserLoginMessage{})
	db.AutoMigrate(&models.Message{})*/
	//db.AutoMigrate(&models.User{}, &models.UserLoginMessage{}, &models.Message{})

	//创建表时添加表后缀
	//db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.User{}, &models.UserLoginMessage{}, &models.Message{})

	//默认路由引擎
	router := gin.Default()

	//前端页面模板解析
	router.LoadHTMLGlob("views/*")

	//路由分离
	routers.ParentRouter(router)

	//启动http服务 默认端口号为127.0.0.1:8080
	router.Run()
}
