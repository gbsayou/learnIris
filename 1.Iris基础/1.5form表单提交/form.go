package main

import "github.com/kataras/iris"

func main() {
	app := iris.Default()

	//定义一个路由，接收一个 form，期待从中获取 message 和 nick 两个字段
	app.Post("/form_post", func(ctx iris.Context) {
		message := ctx.FormValue("message")               //直接中 FormValue 中读取
		nick := ctx.FormValueDefault("nick", "anonymous") //如果没有，则去默认值

		ctx.JSON(iris.Map{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	app.Post("/post", func(ctx iris.Context) {
		id := ctx.URLParam("id")                 // url 中获取
		page := ctx.URLParamDefault("page", "0") // url 中获取，带默认值
		name := ctx.FormValue("name")            //表单中获取
		message := ctx.FormValue("message")      //表单中获取
		// or `ctx.PostValue` for POST, PUT & PATCH-only HTTP Methods.

		app.Logger().Infof("id: %s; page: %s; name: %s; message: %s", id, page, name, message) //记录日志
	})

	app.Run(iris.Addr(":8080"))
}
