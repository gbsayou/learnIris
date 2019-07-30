package main

import (
	"./models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kataras/iris"
)

func main() {

	app := iris.Default()
	app.Logger().SetLevel("debug")

	// app.Use(recover.New()) // 从错误中恢复
	// app.Use(logger.New())  // 记录日志

	db, err := gorm.Open("mysql", "root:123456@/iris?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("数据库连接错误")
	}
	defer db.Close()

	app.Get("/", func(ctx iris.Context) {
		var users []models.User
		result := db.Where(&models.User{Name: "gbs"}).Find(&users)
		ctx.JSON(result)
	})
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed)) // 跑起来
}
